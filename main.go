package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dafalo/Gym-Locator/api"
	"github.com/dafalo/Gym-Locator/models"
	"github.com/dafalo/Gym-Locator/views"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	BindIP = "0.0.0.0"
	Port   = ":8051"
)

func main() {
	u, _ := url.Parse("http://" + BindIP + Port)
	fmt.Printf("Server Started: %v\n", u)

	CreateDB("gym")
	MigrateDB()
	CreateDefaultUser()
	Handlers()

	http.ListenAndServe(Port, nil)
}



func Handlers() {
	fmt.Println("Handlers")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates/"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", views.LoginHandler)
	http.HandleFunc("/GYM", views.GYMHandler)
	http.HandleFunc("/explore", views.ExploreHandler)
	http.HandleFunc("/GymDashboard", views.GymDashboardHandler)
	http.HandleFunc("/GymProfile", views.GymProfileHandler)
	http.HandleFunc("/GymDetails", views.GymDetailsHandler)
	http.HandleFunc("/GymFacility", views.GymFacilityHandler)
	http.HandleFunc("/GymAmenities", views.GymAmenityHandler)
	http.HandleFunc("/GymEquipment", views.GymEquipmentHandler)
	http.HandleFunc("/Login", views.SignHandler)
	http.HandleFunc("/api/", api.APIHandler)
	http.HandleFunc("/logout", views.LogOutHandler)





}

func CreateDB(name string) *sql.DB {
	fmt.Println("Database Created")
	db, err := sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		panic(err)
	}
	db.Close()

	db, err = sql.Open("mysql", "root:a@tcp(127.0.0.1:3306)/"+name)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func MigrateDB() {
	fmt.Println("Database Migrated")
	user := models.User{}
	gym := models.Gym{}
	facility := models.Facility{}
	amenities := models.Amenities{}
	equipment := models.Equipment{}
	


	db := GormDB()
	db.AutoMigrate(&user,&gym,&facility,&amenities,&equipment)
}



func CreateDefaultUser() {

	db := GormDB()

	user := []models.User{}
	db.Find(&user)

	defaultUser := []models.User{
		{
			Username: "admin",
			Password: hashPassword("admin"),
			Name:     "Software Developer",
			Type:     "Administrator",
			
			
		},

		{
			Username: "user",
			Password: hashPassword("user"),
			Name:     "Software Developer",
			Type:     "User",
		
		},

		
	}

	isExisting := false
	for i := range defaultUser {
		isExisting = false

		for _, users := range user {
			if defaultUser[i].Username == users.Username {
				isExisting = true
				break
			}
		}

		if !isExisting {
			fmt.Println("Create Default User")
			db.Save(&defaultUser[i])
		}
	}

}

func hashPassword(pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func GormDB() *gorm.DB {
	dsn := "root:a@tcp(127.0.0.1:3306)/gym?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

