package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	registerEnvironmentRoutes()
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("listening on :%s...\n", port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
