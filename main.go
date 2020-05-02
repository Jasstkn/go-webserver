package main

import (
	"fmt"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"result\": \"OK\"}")
}

func main() {
	http.HandleFunc("/", Root)
	http.HandleFunc("/health", Health)
	log.Fatal(http.ListenAndServe(":8000", nil))
}