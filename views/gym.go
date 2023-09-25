package views

import (
	"net/http"
	"text/template"
)

func GYMHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/gym-register.html"))
	data := map[string]interface{}{}
	data["Title"] = "Main | POS_SYSTEM"
	tmpl.Execute(w, data)
}
