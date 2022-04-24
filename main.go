package main

import (
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleIndex)
	
	log.Println("Starting server on port: :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
