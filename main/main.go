package main

import (
	"fmt"
	"log"
)

func main() {
	server, err := BuildDependencies()
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	port := fmt.Sprintf(":%s", server.config.Port)
	fmt.Printf("Starting server on port %s...\n", port)
	if err := server.Run(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
