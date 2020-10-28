package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	UserID    uint
	EventType string `gorm:"not null" binding:"required,oneof=Office Remote TimeOff Vacation"`
	Note      string
	Date      time.Time `gorm:"not null" binding:"required"`
}

type EventDTO struct {
	EventType string `json:"select-type" binding:"required,oneof=Office Remote TimeOff Vacation"`
	Note      string `json:"note"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

func (event *Event) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&Statistic{}).Where("user_id = ?", event.UserID).UpdateColumn(event.EventType, gorm.Expr(event.EventType+"+ ?", 1))
	return
}

func (event *Event) BeforeUpdate(tx *gorm.DB) (err error) {
	var prev Event
	tx.Where("id = ?", event.ID).First(&prev)
	tx.Model(&Statistic{}).Where("user_id = ?", prev.UserID).Updates(map[string]interface{}{prev.EventType: gorm.Expr(prev.EventType+"- ?", 1)})
	return
}

func (event *Event) AfterUpdate(tx *gorm.DB) (err error) {
	tx.Model(&Statistic{}).Where("user_id = ?", event.UserID).UpdateColumn(event.EventType, gorm.Expr(event.EventType+"+ ?", 1))
	return
}

func (event *Event) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(&Statistic{}).Where("user_id = ?", event.UserID).UpdateColumn(event.EventType, gorm.Expr(event.EventType+"- ?", 1))
	return
}
