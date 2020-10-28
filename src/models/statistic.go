package models

import (
	"gorm.io/gorm"
)

type Statistic struct {
	gorm.Model
	UserID   uint
	Vacation int `gorm:"default:0"`
	TimeOff  int `gorm:"default:0"`
	Remote   int `gorm:"default:0"`
	Office   int `gorm:"default:0"`
}
