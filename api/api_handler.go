package api

import (
	"net/http"
	"strings"
)

// APIHandler !
func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	

	if strings.HasPrefix(r.URL.Path, "login") {
		Login(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "user") {
		CreateUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_user") {
		GetUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_user") {
		EditUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_user") {
		DeleteUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "facility") {
		CreateFacility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_facility") {
		GetFacility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_facility") {
		EditFacility(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_facility") {
		DeleteFacility(w, r)
		return
	}
	
}
