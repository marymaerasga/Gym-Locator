package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateTrainer(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	user := []models.Trainer{}
	db.Find(&user)
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
	test := ""


	defaultUser := []models.Trainer{
		{
			FirstName : fname,
			LastName : lname,
			Email : email,
			Number : phone,
			Address : address,
			DOB : dob,
			Gender: gender,
			Experience : experience,
			Biography : bio,
			GymID : id,
			Password : hashPassword("Gym2023"),
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].FirstName == users.FirstName && defaultUser[i].LastName == users.LastName {
				isExisting = true
				data := map[string]interface{}{
					"status": "Exist",
				}
				ReturnJSON(w, r, data)
				break

				
			}
		}

		if !isExisting {
			db.Save(&defaultUser[i])
			data := map[string]interface{}{
				"status": "New",
			}
			train := models.Trainer{}
			db.Last(&train)
			test = fmt.Sprint(train.ID)

			
			ReturnJSON(w, r, data)
			
		}

		
	}
			var c []map[string]string
			json.Unmarshal([]byte(certifications), &c)
			
			for i := range c {
				certificate := models.Certificate{}
				cert := c[i]["certificate"]
				

				certificate.TrainerID = test
				certificate.Name = cert
				db.Save(&certificate)
				
			}


			var d []map[string]string
			json.Unmarshal([]byte(specialties), &d)
			
			for i := range d {
				special := models.Specialties{}
				cert := d[i]["special"]
			

				special.TrainerID = test
				special.Name = cert
				db.Save(&special)
				
			}

			

	
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
	// certifications := r.FormValue("certifications")
	// specialties := r.FormValue("specialties")
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


