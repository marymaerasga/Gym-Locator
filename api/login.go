package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dafalo/Gym-Locator/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	email := r.FormValue("email")
	username := r.FormValue("username")
	id := r.FormValue("id")
	password := r.FormValue("password")
	position := r.FormValue("position")

	if position == "Owner"{
		user := models.Gym{}
		report := []models.Gym{}
		
		db.Where("email = ?", email).Find(&user)
		db.Where("email = ?", email).Find(&report)
	
		db.Where("id", user.ID).Find(&user)
	
		if CheckPasswordHash(password, user.Password) {
			result := "1"
	
			newSession := uuid.NewString()
	
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "session",
				Value: newSession,
			})
	
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "id",
				Value: fmt.Sprint(user.ID),
			})
			data := map[string]interface{}{
				"status":  "ok",
				"results": result,
				"reports": report,
			}
			ReturnJSON(w, r, data)
		} else {
			result := "0"
			data := map[string]interface{}{
				"status":  "error",
				"results": result,
			}
			ReturnJSON(w, r, data)
		}
		}else if position == "Client"{
			user := models.Client{}
			report := []models.Client{}
			
			db.Where("username = ? && gym_id = ?", username,id).Find(&user)
			db.Where("username = ? && gym_id = ?", username,id).Find(&report)
		
			db.Where("id", user.ID).Find(&user)
		
			if CheckPasswordHash(password, user.Password) {
				result := "3"
		
				newSession := uuid.NewString()
		
				http.SetCookie(w, &http.Cookie{
					Path:  "/",
					Name:  "session",
					Value: newSession,
				})
		
				http.SetCookie(w, &http.Cookie{
					Path:  "/",
					Name:  "client",
					Value: fmt.Sprint(user.ID),
				})
				data := map[string]interface{}{
					"status":  "ok",
					"results": result,
					"reports": report,
				}
				ReturnJSON(w, r, data)
			} else {
				result := "0"
				data := map[string]interface{}{
					"status":  "error",
					"results": result,
				}
				ReturnJSON(w, r, data)
			}
			}else{
		user := models.Trainer{}
		report := []models.Trainer{}
		
		db.Where("email = ?", email).Find(&user)
		db.Where("email = ?", email).Find(&report)
	
		db.Where("id", user.ID).Find(&user)
	
		if CheckPasswordHash(password, user.Password) {
			result := "2"
	
			newSession := uuid.NewString()
	
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "session",
				Value: newSession,
			})
	
			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "id",
				Value: fmt.Sprint(user.GymID),
			})

			http.SetCookie(w, &http.Cookie{
				Path:  "/",
				Name:  "trainer",
				Value: fmt.Sprint(user.ID),
			})
			data := map[string]interface{}{
				"status":  "ok",
				"results": result,
				"reports": report,
			}
			ReturnJSON(w, r, data)
		} else {
			result := "0"
			data := map[string]interface{}{
				"status":  "error",
				"results": result,
			}
			ReturnJSON(w, r, data)
		}
	}



	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}

func CheckPasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}

func GormDB() *gorm.DB {
	dsn := "root:GroupNB2023@tcp(127.0.0.1:3306)/gym?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	w.Write(b)
}



