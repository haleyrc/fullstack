// +build prod

package main

import (
	"net/http"
)

func registerEnvironmentRoutes() {
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/index.html")
	})
	http.Handle("/", http.FileServer(http.Dir("./build")))
}
