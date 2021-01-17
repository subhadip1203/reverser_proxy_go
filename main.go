package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	fmt.Println("server running on : http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello\n")
}
