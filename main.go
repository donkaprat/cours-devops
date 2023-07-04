package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	http.HandleFunc("/toto", func(w http.ResponseWriter, r *http.Request) {
		// Renvoie une erreur 404 (Not Found)
		http.Error(w, "Page introuvable", http.StatusNotFound)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}
