package main

import (
	"fmt"
	"log"
)

func main() {
	server, err := InitDependencies()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := server.Run(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
