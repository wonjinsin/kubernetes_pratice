package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// get env variables
	exampleEnv := os.Getenv("EXAMPLE_ENV")
	if exampleEnv == "" {
		log.Fatal("EXAMPLE_ENV is not set")
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! %s", exampleEnv)
	})
	http.ListenAndServe(":8080", nil)
}
