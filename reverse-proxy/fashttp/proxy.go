package fashttp

import (
	"log"
	"net"
	"net/url"

	"github.com/valyala/fasthttp"
)

// NewReverseProxy ...
func NewReverseProxy(addr string) *ReverseProxy {
	client := &fasthttp.HostClient{
		Addr: addr,
	}

	return &ReverseProxy{
		client: client,
	}
}

type ReverseProxy struct {
	client *fasthttp.HostClient
}

// ServeHTTP ReverseProxy to serve
func (p *ReverseProxy) ServeHTTP(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response

	log.Printf("uri %v", req.URI())


	// prepare request(replace headers and some URL host)
	if clientIP, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {
		req.Header.Add("X-Forwarded-For", clientIP)
		req.Header.Add("X-Forwarded-Host", p.client.Addr)
	}

	// to save all response header
	resHeaders := make(map[string]string)
	res.Header.VisitAll(func(k, v []byte) {

		key := string(k)
		value := string(v)

		if key == "Content-Type" {
			return
		}

		//log.Printf("Header %v : %v", key, value)


		if val, ok := resHeaders[key]; ok {
			resHeaders[key] = val + "," + value
		}
		resHeaders[key] = value
	})

	for _, h := range hopHeaders {
		// if h == "Te" && hv == "trailers" {
		// 	continue
		// }
		req.Header.Del(h)
	}

	//ctx.Logger().Printf("recv a request to proxy to: %s", p.client.Addr)
	if err := p.client.Do(req, res); err != nil {
		ctx.Logger().Printf("could not proxy: %v\n", err)
		return
	}

	// response to client
	for _, h := range hopHeaders {
		res.Header.Del(h)
	}
	for k, v := range resHeaders {
		//if k == "Content-Type" {
		//	v = "text/html"
		//}
		//log.Printf("setting header %v: %v", k, v)
		res.Header.Set(k, v)
	}
}


func ReplaceRequestPath(old, new string) func(ctx *fasthttp.RequestCtx) {

	return func(ctx *fasthttp.RequestCtx) {
		req := &ctx.Request

		u, _ := url.Parse(req.URI().String())
		log.Printf("Raw Path %v", u.RawPath)
		log.Printf("Path %v", u.Path)

		if u.Path ==  old {
			u.Path = new
		}
		req.SetRequestURI(u.RequestURI())
	}
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
