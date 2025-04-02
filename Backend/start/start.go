package ycodrive

import (
	"fmt"
	"net/http"
	"ycodrive/Backend/handler"
)

// StartServer lance le serveur web
func StartServer() {
	// Servir les fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static_css"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	// Routes
	http.HandleFunc("/", ycodrive.HomeHandler)
	http.HandleFunc("/signup", ycodrive.SignupHandler)

	// Lancer le serveur
	port := ":8080"
	fmt.Printf("Serveur démarré sur http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Erreur au démarrage du serveur : %v\n", err)
	}
}
