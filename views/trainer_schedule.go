package views

import (
	"net/http"
	"text/template"
)

func TrainerScheduleHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/trainor/trainor-schedule.html"))
	data := map[string]interface{}{}
	data["Title"] = "Main | POS_SYSTEM"
	tmpl.Execute(w, data)
}
