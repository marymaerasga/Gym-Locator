package models

import "gorm.io/gorm"

type Certificate struct {
	ID       	uint `gorm:"primaryKey"`
	Trainer		 Trainer
	TrainerID	 string
	Name	  	string
	Deleted  gorm.DeletedAt
}


