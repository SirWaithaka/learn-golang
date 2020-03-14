package fasproxy

import (
	"log"
	"net"
	"net/url"
	"strings"

	"github.com/valyala/fasthttp"
)

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// NewReverseProxy ...
func NewReverseProxy(addr *url.URL) *ReverseProxy {
	addrQuery := addr.RawQuery
	director := func(req *fasthttp.Request) {
		reqURL, _ := url.Parse(req.URI().String())

		reqURL.Scheme = addr.Scheme
		reqURL.Host = addr.Host
		reqURL.Path = singleJoiningSlash(addr.Path, reqURL.Path)
		if addrQuery == "" || reqURL.RawQuery == "" {
			reqURL.RawQuery = addrQuery + reqURL.RawQuery
		} else {
			reqURL.RawQuery = addrQuery	+ "&" + reqURL.RawQuery
		}

		if len(req.Header.UserAgent()) == 0 {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	client := &fasthttp.HostClient{
		Addr: addr.Host,
	}


	return &ReverseProxy{
		client: client,
		Director: director,
	}
}

// ReverseProxy reverse handler using fasthttp.HostClient
type ReverseProxy struct {
	client *fasthttp.HostClient
	Director func(*fasthttp.Request)
}

func (p *ReverseProxy) ServeHTTP(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response

	// clone except body stream
	var outreq = new(fasthttp.Request)
	req.CopyTo(outreq)
	if req.Header.ContentLength() == 0 {
		outreq.SetBody(nil)
	}

	p.Director(outreq)


	// to save all response header
	resHeaders := make(map[string]string)
	res.Header.VisitAll(func(k, v []byte) {
		key := string(k)
		value := string(v)
		if val, ok := resHeaders[key]; ok {
			resHeaders[key] = val + "," + value
		}
		resHeaders[key] = value
	})


	for _, h := range hopHeaders {
		hv := string(outreq.Header.Peek(h))
		if hv == "" { continue }
		if h == "Te" && hv == "trailers" {
			continue
		}
		outreq.Header.Del(h)
	}

	// prepare request(replace headers and some URL host)
	if clientIP, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {

		if prior := string(outreq.Header.Peek("X-Forwarded-For")); prior != "" {
			clientIP = strings.Join([]string{prior}, ", ") + ", " + clientIP
		}
		outreq.Header.Set("X-Forwarded-For", clientIP)
		outreq.Header.Add("X-Forwarded-Host", p.client.Addr)
		outreq.Header.Set("Host", p.client.Addr)
	}

	log.Printf("Requests headers are as\n%v", outreq.Header.String())

	// perform request.
	if err := p.client.Do(outreq, res); err != nil {
		ctx.Logger().Printf("could not proxy: %v\n", err)
		return
	}

	// response to client
	for _, h := range hopHeaders {
		res.Header.Del(h)
	}
	for k, v := range resHeaders {
		res.Header.Set(k, v)
	}
}

// SetClient ...
func (p *ReverseProxy) SetClient(addr string) *ReverseProxy {
	p.client.Addr = addr
	return p
}

// Reset ...
func (p *ReverseProxy) Reset() {
	p.client.Addr = ""
}

// Close ...
func (p *ReverseProxy) Close() {
	p.client = nil
	p = nil
}

//func upgradeType(h fasthttp.RequestHeader) string {
//	connHead := h.Peek("Connection")
//	if strings.Contains()
//	if !httpguts.HeaderValuesContainsToken(string(connHead), "Upgrade") {
//		return ""
//	}
//	return strings.ToLower(string(h.Peek("Upgrade")))
//}


func copyResponse(src *fasthttp.Response, dst *fasthttp.Response) {
	src.CopyTo(dst)
}

func copyRequest(src *fasthttp.Request, dst *fasthttp.Request) {
	src.CopyTo(dst)
}

func cloneResponse(src *fasthttp.Response) *fasthttp.Response {
	dst := new(fasthttp.Response)
	copyResponse(src, dst)
	return dst
}

func cloneRequest(src *fasthttp.Request) *fasthttp.Request {
	dst := new(fasthttp.Request)
	copyRequest(src, dst)
	return dst
}

// Hop-by-hop headers. These are removed when sent to the backend.
// As of RFC 7230, hop-by-hop headers are required to appear in the
// Connection header field. These are the headers defined by the
// obsoleted RFC 2616 (section 13.5.1) and are used for backward
// compatibility.
var hopHeaders = []string{
	"Connection",
	"Proxy-Connection", // non-standard but still sent by libcurl and rejected by e.g. google
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",      // canonicalized version of "TE"
	"Trailer", // not Trailers per URL above; https://www.rfc-editor.org/errata_search.php?eid=4522
	"Transfer-Encoding",
	"Upgrade",
}

