package main

import (
	"log"
	"net/http"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "not found", http.StatusMethodNotAllowed)
	}

	http.ServeFile(w, r, "templates/index.html")
}

func main() {
	hub := NewHub()
	go hub.Run()

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(hub, w, r)
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}
