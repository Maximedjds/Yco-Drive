package ycodrive

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	ycodrive "ycodrive/Backend/handler"
)

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
	}
}
