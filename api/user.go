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
	odate := r.FormValue("dfrom")
	dto := r.FormValue("dto")
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
	product.D_to = dto
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
	product := models.Gym{}
	name := r.FormValue("fname")
	lname := r.FormValue("lname")
	password := r.FormValue("password")
	email := r.FormValue("email")
	gname := r.FormValue("name")
	contact := r.FormValue("contact")
	location := r.FormValue("location")
	desc := r.FormValue("description")
	odate := r.FormValue("from")
	dto := r.FormValue("to")
	otime := r.FormValue("time")
	ctime := r.FormValue("time1")
	process := r.FormValue("process")
	
	db.Where("id", id).Find(&product)

	if process == "Profile"{

		product.FirstName = name
		product.LastName = lname
		product.Password = hashPassword(password)
		product.Email = email
	
		db.Save(&product)
	}else{
		product.Name = gname
		product.Contact = contact
		product.Location = location
		product.Description = desc
		product.O_Date = odate
		product.D_to = dto
		product.O_Time = otime
		product.C_Time = ctime

		db.Save(&product)
	}


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
