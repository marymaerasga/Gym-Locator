package models

import "gorm.io/gorm"

type ClassTrainer struct {
	ID       	uint `gorm:"primaryKey"`
	Classes		Classes
	ClassesID	string
	Trainer		Trainer
	TrainerID	 string
	Deleted  gorm.DeletedAt
}


