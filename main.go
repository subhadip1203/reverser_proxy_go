package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var severCount = 0

const (
	SERVER1 = "http://localhost:8000"
	SERVER2 = "http://localhost:8001"
	SERVER3 = "http://localhost:8002"
	PORT    = ":9000"
)

func getProxyURL() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}
	server := servers[severCount]
	severCount++
	// reset the counter and start from the beginning
	if severCount >= len(servers) {
		severCount = 0
	}
	return server
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)
	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)
	// Note that ServeHttp is non blocking & uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func loadBalacer(res http.ResponseWriter, req *http.Request) {
	// Get address of one backend server on which we forward request
	url := getProxyURL()

	// Forward request to original request
	serveReverseProxy(url, res, req)
}

func main() {

	fmt.Println("server running on  http://localhost" + PORT)
	http.HandleFunc("/", loadBalacer)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}
