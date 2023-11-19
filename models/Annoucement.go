package models

import "gorm.io/gorm"

type Annoucement struct {
	ID       uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	Name	 	string
	Deleted  gorm.DeletedAt
}


