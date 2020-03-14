package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"net/url"
	"os"
	proxy "reverse_proxy/fashttp"
)

// A proxy server for fasthttp.


// ProxyHandler ... fasthttp.RequestHandler func
func ProxyHandler(proxyServer *proxy.ReverseProxy) func(*fasthttp.RequestCtx) {
	// all proxy to localhost

	return func(ctx *fasthttp.RequestCtx) {
		proxyServer.ServeHTTP(ctx)
	}
}

func main() {

	if os.Getenv("PROXY_URL") == "" {
		log.Printf("PROXY_URL env not set!")
		os.Exit(1)
	}

	proxyServer := proxy.NewReverseProxy(os.Getenv("PROXY_URL"))

	addr, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	log.Printf("addr scheme %v", addr.Scheme)

	if err := fasthttp.ListenAndServe(":2721", ProxyHandler(proxyServer)); err != nil {
		log.Fatal(err)
	}
}


