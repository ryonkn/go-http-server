package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><body>")
	for key := range r.Header {
		fmt.Fprintln(w, key, ": ", r.Header[key], "</br>")
	}
	fmt.Fprintln(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
