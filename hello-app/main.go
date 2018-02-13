package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "9000"
	if fromEnv := os.Getenv("HTTP_PORT"); fromEnv != "" {
		port = fromEnv
	}

	server := http.NewServeMux()
	server.HandleFunc("/", indexHandler)

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}
