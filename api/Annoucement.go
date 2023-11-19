package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateAnnouncement(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	user := []models.Annoucement{}
	db.Find(&user)
	name := r.FormValue("name")
	id := r.FormValue("id")
	


	defaultUser := []models.Annoucement{
		{
			
			GymID : id,
			Name: name,
		},
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Name == users.Name {
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

func GetAnnouncement(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Annoucement{}
	db.Preload("Gym").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditAnnouncement(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Annoucement{}
	name := r.FormValue("name")

	id := r.FormValue("id")
	
	db.Where("id", c_id).Find(&product)

	product.Name = name
	product.GymID = id
	db.Save(&product)


}

func DeleteAnnouncement(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Annoucement{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}


