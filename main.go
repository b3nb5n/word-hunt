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
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)

	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
