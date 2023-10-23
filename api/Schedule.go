package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateSchedule(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Schedule{}
	
	date := r.FormValue("date")
	to := r.FormValue("to")
	from := r.FormValue("from")
	id := r.FormValue("id")
	t_id := r.FormValue("t_id")
	
	product.Date = date
	product.Otime = to
	product.Ctime = from
	product.GymID = id
	product.TrainerID = t_id
	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetSchedule(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Schedule{}
	db.Preload("Gym").Preload("Trainer").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditSchedule(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Schedule{}
	
	date := r.FormValue("date")
	to := r.FormValue("to")
	from := r.FormValue("from")
	id := r.FormValue("id")
	t_id := r.FormValue("t_id")
	
	db.Where("id", c_id).Find(&product)

	product.Date = date
	product.Otime = to
	product.Ctime = from
	product.GymID = id
	product.TrainerID = t_id
	db.Save(&product)


}

func DeleteSchedule(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Schedule{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


