package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Booking{}
	
	date := r.FormValue("date")
	trainer := r.FormValue("trainer")
	id := r.FormValue("id")
	gid := r.FormValue("t_id")
	
	product.Date = date
	product.GymID = gid
	product.ClientID = id
	product.TrainerID = trainer

	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetBooking(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Booking{}
	db.Preload("Gym").Preload("Trainer").Preload("Client").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditBooking(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Booking{}
	
	date := r.FormValue("date")
	trainer := r.FormValue("trainer")
	id := r.FormValue("id")
	gid := r.FormValue("t_id")
	
	db.Where("id", c_id).Find(&product)

	product.Date = date
	product.GymID = gid
	product.ClientID = id
	product.TrainerID = trainer
	db.Save(&product)


}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Booking{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


