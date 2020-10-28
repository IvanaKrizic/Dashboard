package repositories

import (
	"github.com/IvanaKrizic/Dashboard/src/config"
	"github.com/IvanaKrizic/Dashboard/src/models"
)

func DeleteAuth(token string) error {
	var auth models.AuthData
	if err := config.DB.Where("token = ?", token).Delete(&auth).Error; err != nil {
		return err
	}
	return nil
}

func CreateAuth(ad *models.AuthData) error {
	if err := config.DB.Create(ad).Error; err != nil {
		return err
	}
	return nil
}

func GetAuth(token string) error {
	var auth models.AuthData
	if err := config.DB.Where("token = ?", token).First(&auth).Error; err != nil {
		return err
	}
	return nil
}
