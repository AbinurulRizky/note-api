package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username 		string		`gorm:"not null" json:"username"`
	Email			string		`gorm:"unique;not null" json:"email"`
	// Password Hashing should be handled before daving to database
	PasswordHash	string		`gorm:"not null" json:"-"`
}