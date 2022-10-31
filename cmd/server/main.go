package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"
)

const (
	portEnvVar = "PORT"
)

func main() {
	rawPort := os.Getenv(portEnvVar)
	port, ok := strconv.ParseInt(rawPort)
	serverAddr := fmt.Sprintf(":%d", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(serverAddr, mux)
}
