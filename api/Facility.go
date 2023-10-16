package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Facility{}
	name := r.FormValue("name")
	description := r.FormValue("description")
	id := r.FormValue("id")
	
	product.Name = name
	product.Description = description
	product.GymID = id
	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Facility{}
	db.Preload("Gym").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Facility{}
	name := r.FormValue("name")
	description := r.FormValue("description")
	id := r.FormValue("id")
	
	db.Where("id", c_id).Find(&product)

	product.Name = name
	product.Description = description
	product.GymID = id
	db.Save(&product)


}

func DeleteFacility(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Facility{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


