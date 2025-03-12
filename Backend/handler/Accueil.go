package ycodrive

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Server struct {
	templates *template.Template
}

var tp1 *template.Template

func (s *Server) AccueilHandler(w http.ResponseWriter, r *http.Request) {
	// Exécuter le template nommé "accueil.html"
	err := s.templates.ExecuteTemplate(w, "accueil.html", nil)
	if err != nil {
		log.Printf("Erreur lors du rendu du template: %v", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
}

func NewServer(templateDir string) (*Server, error) {
	// Charger tous les templates du répertoire
	tmpl, err := template.ParseGlob(filepath.Join(templateDir, "*.html"))
	if err != nil {
		return nil, err
	}
	return &Server{templates: tmpl}, nil
}
