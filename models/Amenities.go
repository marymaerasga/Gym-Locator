package models

import "gorm.io/gorm"

type Amenities struct {
	ID       	uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	Name	  	string
	Deleted  gorm.DeletedAt
}


