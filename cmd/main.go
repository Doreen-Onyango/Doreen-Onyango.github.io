package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	wrappedRouter, err := Setup()
	if err != nil {
		log.Fatalf("Setup failed: %v", err)
	}
	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedRouter,
	}
	fmt.Println("listening to http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
