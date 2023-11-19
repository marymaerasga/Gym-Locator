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

	if strings.HasPrefix(r.URL.Path, "amenities") {
		CreateAmenities(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_amenities") {
		GetAmenities(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_amenities") {
		EditAmenities(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_amenities") {
		DeleteAmenities(w, r)
		return
	}


	if strings.HasPrefix(r.URL.Path, "equipment") {
		CreateEquipment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_equipment") {
		GetEquipment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_equipment") {
		EditEquipment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_equipment") {
		DeleteEquipment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "trainer") {
		CreateTrainer(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_trainer") {
		GetTrainer(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_trainer") {
		EditTrainer(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_trainer") {
		DeleteTrainer(w, r)
		return
	}


	if strings.HasPrefix(r.URL.Path, "class") {
		CreateClass(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_class") {
		GetClass(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_class") {
		EditClass(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_class") {
		DeleteClass(w, r)
		return
	}


	if strings.HasPrefix(r.URL.Path, "schedule") {
		CreateSchedule(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_schedule") {
		GetSchedule(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_schedule") {
		EditSchedule(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_schedule") {
		DeleteSchedule(w, r)
		return
	}


	if strings.HasPrefix(r.URL.Path, "client") {
		CreateClient(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_client") {
		GetClient(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "edit_client") {
		EditClient(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "delete_client") {
		DeleteClient(w, r)
		return
	}



	if strings.HasPrefix(r.URL.Path, "booking") {
		CreateBooking(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_booking") {
		GetBooking(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_booking") {
		EditBooking(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_booking") {
		DeleteBooking(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "plan") {
		CreatePlan(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_plan") {
		GetPlan(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_plan") {
		EditPlan(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_plan") {
		DeletePlan(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "payment") {
		EditPayment(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "ment") {
		CreateAnnouncement(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_ment") {
		GetAnnouncement(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_ment") {
		EditAnnouncement(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_ment") {
		DeleteAnnouncement(w, r)
		return
	}
	
	
	
	
}
