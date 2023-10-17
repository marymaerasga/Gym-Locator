package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateAmenities(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Amenities{}
	name := r.FormValue("name")
	desc := r.FormValue("description")
	id := r.FormValue("id")
	
	product.Name = name
	product.Description = desc
	product.GymID = id
	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetAmenities(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Amenities{}
	db.Preload("Gym").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditAmenities(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Amenities{}
	name := r.FormValue("name")
	desc := r.FormValue("description")
	id := r.FormValue("id")
	
	db.Where("id", c_id).Find(&product)

	product.Name = name
	product.Description = desc
	product.GymID = id
	db.Save(&product)


}

func DeleteAmenities(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Amenities{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


