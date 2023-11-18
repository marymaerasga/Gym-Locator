package models

import "gorm.io/gorm"

type Specialties struct {
	ID       	uint `gorm:"primaryKey"`
	Trainer		 Trainer
	TrainerID	 string
	Name	  	string
	Deleted  gorm.DeletedAt
}


