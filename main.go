package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health-check", HealthCheckHandler)
	fmt.Println("Server Web started on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"alive": true}`)
}