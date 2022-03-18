package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	http.HandleFunc("/bar",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
	http.HandleFunc("/hello",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "🍫 Welcome to the Matrix 👍🏼🎉")
		})

	http.ListenAndServe(":8080", nil)
}
