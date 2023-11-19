package models

import "gorm.io/gorm"

type Plan struct {
	ID       uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	Name	 	string
	Duration  	string
	Amount	  	string
	Deleted  gorm.DeletedAt
}


