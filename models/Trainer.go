package models

import "gorm.io/gorm"

type Trainer struct {
	ID       uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	FirstName 	string
	LastName  	string
	Email	  	string
	Number		string
	Address		string
	DOB			string
	Gender		string
	Certifications		string
	Specialty	string
	Experience	string
	Biography	string
	Deleted  gorm.DeletedAt
}


