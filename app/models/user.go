package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Model
	Username       string `gorm:"type:varchar(20);unique_index"  json:"username"`
	Name           string `gorm:"size:255"  json:"name"`
	Bio            string `gorm:"size:255"  json:"bio"`
	Email          string `gorm:"type:varchar(100);unique_index"  json:"email"`
	HashedPassword []byte `json:"-"`
	Active         bool   `json:"active"`
	IsAdmin        bool   `json:"isAdmin"`
	FileName       string `json:"fileName"`
}

// SetNewPassword set a new hashsed password to user
func (user *User) SetNewPassword(passwordString string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	user.HashedPassword = bcryptPassword
}
