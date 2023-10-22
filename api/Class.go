package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateClass(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Classes{}
	name := r.FormValue("name")
	instructor := r.FormValue("instructor")
	date := r.FormValue("date")
	duration := r.FormValue("duration")
	capacity := r.FormValue("capacity")
	description := r.FormValue("description")
	category := r.FormValue("category")
	equipment := r.FormValue("equipment")
	id := r.FormValue("id")
	
	product.Name = name
	product.TrainerID = fmt.Sprint(instructor)
	product.Date = date
	product.Duration = duration
	product.Capacity = capacity
	product.Description = description
	product.Category = category
	product.EquipmentID = fmt.Sprint(equipment)
	product.GymID = id
	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetClass(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Classes{}
	db.Preload("Gym").Preload("Trainer").Preload("Equipment").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditClass(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Classes{}
	name := r.FormValue("name")
	instructor := r.FormValue("instructor")
	date := r.FormValue("date")
	duration := r.FormValue("duration")
	capacity := r.FormValue("capacity")
	description := r.FormValue("description")
	category := r.FormValue("category")
	equipment := r.FormValue("equipment")
	id := r.FormValue("id")
	
	
	db.Where("id", c_id).Find(&product)

	product.Name = name
	product.TrainerID = fmt.Sprint(instructor)
	product.Date = date
	product.Duration = duration
	product.Capacity = capacity
	product.Description = description
	product.Category = category
	product.EquipmentID = fmt.Sprint(equipment)
	product.GymID = id
	db.Save(&product)


}

func DeleteClass(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Classes{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


