package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Server started")

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Write([]byte(fmt.Sprintf("Welcome to the start page, %s", name)))
	})

	http.ListenAndServe(":8080", nil)
}
