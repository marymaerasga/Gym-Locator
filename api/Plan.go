package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreatePlan(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	user := []models.Plan{}
	db.Find(&user)
	name := r.FormValue("name")
	duration := r.FormValue("duration")
	amount := r.FormValue("amount")
	id := r.FormValue("id")


	defaultUser := []models.Plan{
		{
			
			GymID : id,
			Name: name,
			Duration: duration,
			Amount: amount,
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Name == users.Name && defaultUser[i].GymID == users.GymID  {
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
			ReturnJSON(w, r, data)
		}

	}
			
}

func GetPlan(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Plan{}
	db.Preload("Gym").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditPlan(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Plan{}
	name := r.FormValue("name")
	duration := r.FormValue("duration")
	amount := r.FormValue("amount")
	id := r.FormValue("id")
	
	db.Where("id", c_id).Find(&product)

	product.Name = name
	product.Duration = duration
	product.Amount = amount
	product.GymID = id
	db.Save(&product)


}

func DeletePlan(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Plan{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


