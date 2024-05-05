package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Madhunisha K")
	})
	http.HandleFunc("/regnum", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211SE123")
	})
	http.HandleFunc("/dep", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ISE")
	})
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Skyblue")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	fmt.Printf("Server running (port=8082), route: http://localhost:8082/\n")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
