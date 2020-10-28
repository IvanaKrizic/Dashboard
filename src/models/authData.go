package models

type AuthData struct {
	ID    uint `gorm:"primaryKey"`
	Token string
}
