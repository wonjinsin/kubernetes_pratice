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

	exampleEnv2 := os.Getenv("EXAMPLE_ENV2")
	if exampleEnv2 == "" {
		log.Fatal("EXAMPLE_ENV2 is not set")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY is not set")
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! %s %s", exampleEnv, exampleEnv2)
		fmt.Fprintf(w, "Secret is %s", secretKey)
	})
	http.ListenAndServe(":8080", nil)
}
