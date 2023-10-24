package models

import "gorm.io/gorm"

type Booking struct {
	ID       uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	Client		 Client
	ClientID	 string
	Trainer		 Trainer
	TrainerID	 string
	Date		 string
	Deleted  gorm.DeletedAt
}


