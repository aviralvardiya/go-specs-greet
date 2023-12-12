package main

import (
	"log"
	"net/http"
	"go-specs-greet/adapters/httpserver"
)

func main() {
	mux := httpserver.GetServerMux()

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
