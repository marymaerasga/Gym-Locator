package models

import "gorm.io/gorm"

type Equipment struct {
	ID       	uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	Name	  	string
	Description string
	Deleted  gorm.DeletedAt
}


