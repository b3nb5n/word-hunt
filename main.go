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

	log.Printf("listening on port %s\n\n", port)

	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
