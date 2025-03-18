package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	_ = tmpl.Execute(w, nil)
}
