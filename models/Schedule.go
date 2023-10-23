package models

import "gorm.io/gorm"

type Schedule struct {
	ID       uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	Trainer		Trainer
	TrainerID   string
	Date		string
	Otime		string
	Ctime		string
	Deleted  gorm.DeletedAt
}


