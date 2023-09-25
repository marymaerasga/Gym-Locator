package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Gym-Locator/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	product := models.Gym{}
	name := r.FormValue("name")
	lname := r.FormValue("lname")
	password := r.FormValue("password")
	email := r.FormValue("email")
	gname := r.FormValue("gname")
	contact := r.FormValue("contact")
	location := r.FormValue("location")
	desc := r.FormValue("desc")
	odate := r.FormValue("odate")
	otime := r.FormValue("otime")
	ctime := r.FormValue("ctime")
	terms := r.FormValue("terms")
	policy := r.FormValue("policy")
	


	product.FirstName = name
	product.LastName = lname
	product.Password = hashPassword(password)
	product.Email = email
	product.Name = gname
	product.Contact = contact
	product.Location = location
	product.Description = desc
	product.O_Date = odate
	product.O_Time = otime
	product.C_Time = ctime
	product.Terms = terms
	product.Policy = policy
	product.Type = "Test"
	
	

	db.Save(&product)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	item := []models.Gym{}
	db.Find(&item)

	data := map[string]interface{}{
		"status": "ok",
		"item":   item,
	}
	ReturnJSON(w, r, data)
	sqlDB, _ := db.DB()
	sqlDB.Close()
}


func EditUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	product := models.User{}
	name := r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	number := r.FormValue("number")
	
	db.Where("id", id).Find(&product)

	product.Name = name
	product.Username = username
	product.Number = number

	if (password == ""){

	}else{
		product.Password = hashPassword(password)
	}
	db.Save(&product)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	item := models.User{}
	db.Where("id", id).Statement.Delete(&item)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}
