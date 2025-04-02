package ycodrive

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// User représente les données d'un utilisateur
type User struct {
	Email    string
	Password string
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Charger le template HTML
	tmpl, err := template.ParseFiles(filepath.Join("templates", "signup.html"))
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// Afficher le formulaire d'inscription
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Erreur d'affichage", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		// Traiter le formulaire soumis
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur de traitement du formulaire", http.StatusBadRequest)
			return
		}

		user := User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		// Ici, tu pourrais ajouter la logique pour sauvegarder l'utilisateur
		// dans une base de données

		// Réponse de succès
		successMsg := fmt.Sprintf("Utilisateur %s inscrit avec succès", user.Email)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, successMsg)

	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Erreur d'affichage", http.StatusInternalServerError)
		return
	}
}
