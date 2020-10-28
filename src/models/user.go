package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"not null" binding:"required"`
	LastName  string `json:"lastName" gorm:"not null" binding:"required"`
	Email     string `json:"email" gorm:"not null" binding:"required"`
	Password  string `json:"password" gorm:"not null" binding:"required"`
	Statistic *Statistic
	Events    []*Event
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var stat Statistic
	u.Statistic = &stat

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	u.Password = string(hashedPassword)

	if err != nil {
		return err
	}
	return nil
}
