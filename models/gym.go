package models

import "gorm.io/gorm"

type Gym struct {
	ID       uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Email	  string
	Password  string
	Name		string
	Contact		string
	Location	string
	Lat			string
	Long		string
	Description	string
	O_Date		string
	D_to		string
	O_Time		string
	C_Time		string
	Type		string
	Terms		string
	Policy		string
	Deleted  gorm.DeletedAt
}


