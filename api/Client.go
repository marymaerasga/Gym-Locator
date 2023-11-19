package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dafalo/Gym-Locator/models"
)

func CreateClient(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Client{}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	gender := r.FormValue("gender")
	age := r.FormValue("age")
	member := r.FormValue("member")
	username := r.FormValue("username")
	password := r.FormValue("password")
	payment := r.FormValue("payment")
	amount := r.FormValue("amount")
	id := r.FormValue("id")
	now := time.Now()
	date := now.Format("2006-01-02")
	
	product.FirstName = fname
	product.LastName = lname
	product.Number = phone
	product.Address = address
	product.Gender = gender
	product.Age = age
	product.Member = member
	product.Username = username
	product.Password = hashPassword(password)
	product.Date = date
	product.GymID = id
	product.Payment = payment
	product.Price = amount
	product.Status = "Pending"
	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetClient(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Client{}
	db.Preload("Gym").Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditClient(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	product := models.Client{}
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	phone := r.FormValue("number")
	address := r.FormValue("address")
	age := r.FormValue("age")
	gid := r.FormValue("gid")
	
	db.Where("id", id).Find(&product)

	product.FirstName = fname
	product.LastName = lname
	product.Number = phone
	product.Address = address
	product.Age = age
	product.GymID = gid
	db.Save(&product)


}

func EditPayment(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	c_id, _ := strconv.Atoi(r.FormValue("c_id"))
	product := models.Client{}
	amount := r.FormValue("amount")


	db.Where("id", c_id).Find(&product)

	product.Payment = amount
	db.Save(&product)

}

func DeleteClient(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.Client{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}





