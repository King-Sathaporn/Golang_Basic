package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Profile Page, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Login Page, %s, %s", r.URL.Query().Get("username"), r.URL.Query().Get("password"))
		//get the query string from the request and print it out.

	})

	http.ListenAndServe(":8080", nil) // ListenAndServe starts an HTTP server with a given address and handler.
}
