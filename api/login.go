package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

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
	// id := r.FormValue("id")
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
			
			db.Where("username = ? ", username).Find(&user)
			db.Where("username = ? ", username).Find(&report)
		
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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
   b := make([]byte, length)
   for i := range b {
      b[i] = charset[seededRand.Intn(len(charset))]
   }
   return string(b)
}


func Email(w http.ResponseWriter, r *http.Request) {
	db := GormDB()
	email := r.FormValue("email")
	position := r.FormValue("position")
	password := generateRandomString(10)

	if position == "Owner"{
		user := models.Gym{}
		report := []models.Gym{}
		
		db.Where("email = ?", email).Find(&user)
		db.Where("email = ?", email).Find(&report)
	
		length := len(report)

		if length < 1{
			result := "0"
			data := map[string]interface{}{
				"status":  "error",
				"results": result,
			}
			ReturnJSON(w, r, data)

		}else{
			from := "dffalo.amg.pps@gmail.com"
		pass := "hsejfkwaabbkiqrr"
		to := user.Email
	
		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Change Password\n\n" +
		   "Good day,\n\n Your New Password to Gym-Locator App is" + " " + password + "\n\n Regards \n\n Gym-Locator App "
	
		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))
	
		if err != nil {
			fmt.Printf("smtp error: %s", err)
			return
		}
		fmt.Println("Successfully sended to " + to)

		fmt.Println("password",password)

		user.Password = hashPassword(password)
		db.Save(&user)
			result := "1"
			data := map[string]interface{}{
				"status":  "ok",
				"results": result,
				"reports": report,
			}
			ReturnJSON(w, r, data)
		}
	}else{
		user := models.Trainer{}
		report := []models.Trainer{}
		
		db.Where("email = ?", email).Find(&user)
		db.Where("email = ?", email).Find(&report)

		length := len(report)

		if length < 1{
			result := "0"
			data := map[string]interface{}{
				"status":  "error",
				"results": result,
			}
			ReturnJSON(w, r, data)

		}else{
			from := "dffalo.amg.pps@gmail.com"
		pass := "hsejfkwaabbkiqrr"
		to := user.Email
	
		msg := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: Change Password\n\n" +
		   "Good day,\n\n Your New Password to Gym-Locator App is" + " " + password + "\n\n Regards \n\n Gym-Locator App "
	
		err := smtp.SendMail("smtp.gmail.com:587",
			smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
			from, []string{to}, []byte(msg))
	
		if err != nil {
			fmt.Printf("smtp error: %s", err)
			return
		}
		fmt.Println("Successfully sended to " + to)

		fmt.Println("password",password)

		user.Password = hashPassword(password)
		db.Save(&user)
			result := "2"
			data := map[string]interface{}{
				"status":  "ok",
				"results": result,
				"reports": report,
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
	dsn := "root:a@tcp(127.0.0.1:3306)/gym?charset=utf8mb4&parseTime=True&loc=Local"
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



