package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"net/url"
	"os"
	"reverse_proxy/fasproxy"
)

var proxyServer *fasproxy.ReverseProxy


// ProxyHandler ... fasthttp.RequestHandler func
func FasProxyHandle(ctx *fasthttp.RequestCtx) {
	// all proxy to localhost
	proxyServer.ServeHTTP(ctx)
}

func main() {
	//proxyURL, err := url.Parse("http://localhost:1338")
	proxyURL, err := url.Parse("https://splash.youtise.com")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	proxyServer = fasproxy.NewReverseProxy(proxyURL)

	log.Printf("Host name %v\n", proxyURL.Host)

	if err := fasthttp.ListenAndServe(":2720", FasProxyHandle); err != nil {
		log.Fatal(err)
	}
}
