package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8000",
	})

	fmt.Println("server running on  http://localhost:9000")
	if err := http.ListenAndServe(":9000", proxy); err != nil {
		panic(err)
	}
}
