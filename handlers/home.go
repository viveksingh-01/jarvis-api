package handlers

import (
	"html/template"
	"net/http"
)

var homeTemplate = template.Must(template.ParseFiles("templates/home.html"))

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	homeTemplate.Execute(w, nil)
}
