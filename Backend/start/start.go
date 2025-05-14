package ycodrive

import (
	"fmt"
<<<<<<< HEAD
	"net/http"
	"ycodrive/Backend/handler"
)

// StartServer lance le serveur web
func StartServer() {
	// Servir les fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static_css"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))
}
	// Routes
	http.HandleFunc("/", ycodrive.HomeHandler)
	http.HandleFunc("/signup", ycodrive.SignupHandler)

	// Lancer le serveur
	port := ":8080"
	fmt.Printf("Serveur démarré sur http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Erreur au démarrage du serveur : %v\n", err)
=======
	"html/template"
	"log"
	"net/http"
	ycodrive "ycodrive/Backend/handler"
	}

var tp1 *template.Template

func Start() {
	// Initialisation du serveur avec le répertoire des templates
	server, err := ycodrive.NewServer("../frontend/templates")
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates: %v", err)
	}

	// Gestion des fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("../frontend/static_css"))))

	// Routes
	http.HandleFunc("/accueil", server.AccueilHandler)

	fmt.Println("Serveur lancé sur: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erreur serveur: %v", err)
>>>>>>> a7db7c616b8cb6590112a949731ce72e42e70698
	}
}
