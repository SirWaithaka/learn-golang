package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"net"
	"net/url"
	"os"
)


var hopHeaders = []string{
	"Connection",
	"Proxy-Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailer",
	"Transfer-Encoding",
	"Upgrade",
}


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
	}

	req.Header.Add("X-Forwarded-Host", p.client.Addr)
	req.Header.Set("Host", p.client.Addr)

	// to save all response header
	resHeaders := make(map[string]string)
	res.Header.VisitAll(func(k, v []byte) {

		key := string(k)
		value := string(v)

		if key == "Content-Type" {
			return
		}

		if val, ok := resHeaders[key]; ok {
			resHeaders[key] = val + "," + value
		}
		resHeaders[key] = value
	})

	for _, h := range hopHeaders {
		req.Header.Del(h)
	}

	if err := p.client.Do(req, res); err != nil {
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

// ProxyHandler ... fasthttp.RequestHandler func
func Handler(client *ReverseProxy) func(*fasthttp.RequestCtx) {
	// all proxy to localhost

	return func(ctx *fasthttp.RequestCtx) {
		client.ServeHTTP(ctx)
	}
}

func main() {

	if os.Getenv("PROXY_URL") == "" {
		log.Printf("PROXY_URL env not set!")
		os.Exit(1)
	}

	proxyClient := NewReverseProxy(os.Getenv("PROXY_URL"))

	addr, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	log.Printf("addr scheme %v", addr.Scheme)

	if err := fasthttp.ListenAndServe(":2721", Handler(proxyClient)); err != nil {
		log.Fatal(err)
	}
}

