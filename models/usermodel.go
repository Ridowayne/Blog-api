package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Phone     int64  `gorm:"unique"`
	Password  string
	Role string `gorm:"type:enum('admin', 'user', 'guest');default:'user"`
}