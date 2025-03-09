package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	/*
		This function handles the home page and renders the index page.
	*/
	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tmpl.Execute(w, nil)
}
