package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Equipment{}
	name := r.FormValue("name")
	id := r.FormValue("id")
	
	product.Name = name
	product.GymID = id
	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Equipment{}
	db.Preload("Gym").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Equipment{}
	name := r.FormValue("name")
	id := r.FormValue("id")
	
	db.Where("id", c_id).Find(&product)

	product.Name = name
	product.GymID = id
	db.Save(&product)


}

func DeleteEquipment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Equipment{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


