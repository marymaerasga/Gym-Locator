package models

import "gorm.io/gorm"

type Client struct {
	ID       uint `gorm:"primaryKey"`
	Gym		 	Gym
	GymID	 	string
	FirstName 	string
	LastName  	string
	Number		string
	Address		string
	Gender		string
	Age			string
	Member		string
	Price		string
	Username	string
	Password	string
	Payment		string
	Date 		string
	Status		string
	Deleted  gorm.DeletedAt
}


