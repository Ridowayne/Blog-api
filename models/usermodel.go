package models

import (
	"errors"
	"unicode"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Phone     string  `gorm:"not null;unique;size:10"`
	Password  string
	Role string `gorm:"default:'user'"`
}

func (u *User) Validate() error {
    if len(u.Phone) != 10 {
        return errors.New("phone number must be 10 digits long")
    }
    for _, r := range u.Phone {
        if !unicode.IsDigit(r) {
            return errors.New("phone number can only contain numeric characters")
        }
    }
    return nil
}
