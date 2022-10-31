package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	portEnvVar = "PORT"
)

func main() {
	rawPort := os.Getenv(portEnvVar)
	port, err := strconv.Atoi(rawPort)
	if err != nil {
		log.Fatalf("failed to parse port %s into num: %v", rawPort, err)
	}
	serverAddr := fmt.Sprintf(":%d", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Printf("starting server at address %q\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("listen and serve exited with non-nil err: %v", err)
	}
}
