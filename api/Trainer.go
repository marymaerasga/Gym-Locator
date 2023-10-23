package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateTrainer(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Trainer{}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	dob := r.FormValue("dob")
	gender := r.FormValue("gender")
	certifications := r.FormValue("certifications")
	specialties := r.FormValue("specialties")
	experience := r.FormValue("experience")
	bio := r.FormValue("bio")
	id := r.FormValue("id")
	
	product.FirstName = fname
	product.LastName = lname
	product.Email = email
	product.Number = phone
	product.Address = address
	product.DOB = dob
	product.Gender = gender
	product.Certifications = certifications
	product.Specialty = specialties
	product.Experience = experience
	product.Biography = bio
	product.GymID = id
	product.Password = hashPassword("Gym2023")
	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetTrainer(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Trainer{}
	db.Preload("Gym").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditTrainer(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Trainer{}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	dob := r.FormValue("dob")
	gender := r.FormValue("gender")
	certifications := r.FormValue("certifications")
	specialties := r.FormValue("specialties")
	experience := r.FormValue("experience")
	bio := r.FormValue("bio")
	id := r.FormValue("id")
	
	db.Where("id", c_id).Find(&product)

	product.FirstName = fname
	product.LastName = lname
	product.Email = email
	product.Number = phone
	product.Address = address
	product.DOB = dob
	product.Gender = gender
	product.Certifications = certifications
	product.Specialty = specialties
	product.Experience = experience
	product.Biography = bio
	product.GymID = id
	db.Save(&product)


}

func DeleteTrainer(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Trainer{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


