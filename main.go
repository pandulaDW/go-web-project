package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice
// as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on port 4000...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
