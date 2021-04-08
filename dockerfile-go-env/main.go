package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		env := os.Getenv("ENV")
		name := os.Getenv("NAME")
		fmt.Fprintf(w, fmt.Sprintf("Hello %s from %s", env, name))
	})

	http.ListenAndServe(":8080", nil)
}
