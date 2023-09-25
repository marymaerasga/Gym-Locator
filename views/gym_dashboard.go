package views

import (
	"net/http"
	"text/template"
)

func GymDashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/gym-owner/dashboard.html"))
	data := map[string]interface{}{}
	data["Title"] = "Main | POS_SYSTEM"
	tmpl.Execute(w, data)
}
