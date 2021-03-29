package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("straring server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", requestHandler)
	log.Printf("listening on port %s\n\n", port)

	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
