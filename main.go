package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", handler)
	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
