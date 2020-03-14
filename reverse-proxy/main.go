package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

// Get env var or default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Get the port to listen on
func getListenAddress() string {
	port := getEnv("PORT", "1338")
	return ":" + port
}

// Log the typeform payload and redirect url
func logRequestPayload(req *http.Request) {
	if strings.Contains(req.URL.String(), "_next") {
		return
	}
	if strings.Contains(req.URL.String(), "/cdn-cgi") {
		return
	}

	log.Printf("%s : %s from %s\n", req.Method, req.URL, req.RemoteAddr)
}

// Log the env variables required for a reverse proxy
func logSetup() {

	log.Printf("Server will run on: %s\n", getListenAddress())
}

/*
	Reverse Proxy Logic
*/

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	proxyURL, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	log.Println(req.Header.Get(""))

	// Update the headers to allow for SSL redirection
	req.URL.Host = proxyURL.Host
	req.URL.Scheme = proxyURL.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = proxyURL.Host

	log.Printf("host %v", req.Header.Get("Host"))
	log.Printf("request headers %v", req.Header)

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {

	proxyURL := os.Getenv("PROXY_URL")
	if proxyURL == "" {
		proxyURL = "https://splash.youtise.com"
	}


	logRequestPayload(req)

	serveReverseProxy(proxyURL, res, req)
}

func main() {
	// Log setup values
	logSetup()

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
