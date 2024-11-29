package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	text := os.Getenv("TEXT")
	if text == "" {
		text = "Hello, World!"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, text)
	})

	fmt.Printf("Serving '%s' on port %s\n", text, port)
	http.ListenAndServe(":"+port, nil)
}
