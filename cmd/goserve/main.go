package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	host := flag.String("host", "0.0.0.0", "Host to bind the server")
	port := flag.String("port", "8080", "Port to run the server on")
	root := flag.String("root", ".", "Root directory to serve")
	verbose := flag.Bool("v", false, "Enable verbose logging")

	addr := host + ":" + port

	flag.Parse()

	if _, err := os.Stat(*root); os.IsNotExist(err) {
		log.Fatalf("Root directory does not exist: %s", *root)
	}

	if *verbose {
		fmt.Printf("Serving root directory: %s\n", *root)
	}

	http.HandleFunc("/", clientHandler)
	log.Printf("HTTP server runnign on %s:%s", *host, *port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
