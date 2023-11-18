package api

import (
	"encoding/json"
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
	product.Date = date
	product.Duration = duration
	product.Capacity = capacity
	product.Description = description
	product.Category = category
	product.EquipmentID = fmt.Sprint(equipment)
	product.GymID = id
	db.Save(&product)

	var c []map[string]string
	json.Unmarshal([]byte(instructor), &c)

			
	for i := range c {
		println(c[i]["trainer"])
		certificate := models.ClassTrainer{}
		cert := c[i]["trainer"]
		

		certificate.TrainerID = cert
		certificate.ClassesID = fmt.Sprint(product.ID)
		db.Save(&certificate)
		
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetClass(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Classes{}
	db.Preload("Gym").Preload("Equipment").Find(&item)

	class := []models.ClassTrainer{}
	db.Preload("Classes").Preload("Trainer").Find(&class)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
		"class" : class,
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
	// instructor := r.FormValue("instructor")
	date := r.FormValue("date")
	duration := r.FormValue("duration")
	capacity := r.FormValue("capacity")
	description := r.FormValue("description")
	category := r.FormValue("category")
	equipment := r.FormValue("equipment")
	id := r.FormValue("id")
	
	
	db.Where("id", c_id).Find(&product)

	product.Name = name
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


