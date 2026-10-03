package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	handler := http.FileServer(http.Dir("."))
	mux.Handle("/", handler)
	http.ListenAndServe(":8080", mux)
}
