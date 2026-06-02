package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func ExecutarTemplate(w http.ResponseWriter, nome string, dados interface{}) {
	err := templates.ExecuteTemplate(w, nome, dados)
	if err != nil {
		http.Error(w, "Erro ao carregar a página", http.StatusInternalServerError)
	}
}
