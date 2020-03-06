package main

// A proxy server for fasthttp.

import (
	"log"

	"github.com/valyala/fasthttp"
	proxy "reverse_proxy/fashttp"
)

var (
	proxyServer = proxy.NewReverseProxy("localhost:3000")
)

// ProxyHandler ... fasthttp.RequestHandler func
func ProxyHandler(ctx *fasthttp.RequestCtx) {
	// all proxy to localhost
	proxyServer.ServeHTTP(ctx)
}

func main() {
	if err := fasthttp.ListenAndServe(":2720", ProxyHandler); err != nil {
		log.Fatal(err)
	}
}