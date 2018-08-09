package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><body>")
	fmt.Fprintln(w, "<h1>", os.Getenv("HOSTNAME"), "</h1>")

	fmt.Fprintln(w, "<h2>Request Headers</h2>")
	for key := range r.Header {
		fmt.Fprintln(w, key, ": ", r.Header[key], "<br>")
	}

	fmt.Fprintln(w, "<h2>Environments</h2>")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Fprintln(w, pair[0], ": ", pair[1], "<br>")
	}

	fmt.Fprintln(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
