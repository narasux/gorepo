package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		message := []byte("Hello from " + hostname)
		_, _ = w.Write(message)
	})

	fmt.Println("Starting server on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}
