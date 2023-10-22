package models

import "gorm.io/gorm"

type Classes struct {
	ID       uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	Trainer		Trainer
	TrainerID	string
	Name	 	string
	Date	  	string
	Duration	string
	Capacity	string
	Description	string
	Category	string
	Equipment	Equipment
	EquipmentID	string
	Deleted  gorm.DeletedAt
}


