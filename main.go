package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	handler := http.FileServer(http.Dir("."))

	mux.Handle("/", handler)

	log.Fatal(srv.ListenAndServe())


}
