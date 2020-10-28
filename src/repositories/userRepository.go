package repositories

import (
	"github.com/IvanaKrizic/Dashboard/src/config"
	"github.com/IvanaKrizic/Dashboard/src/models"
)

func GetAllUsers(users *[]models.User) (err error) {
	if err = config.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *models.User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *models.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(user *models.User, email string) (err error) {
	if err = config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, id string) (err error) {
	config.DB.Save(user)
	return nil
}

func DeleteUser(user *models.User, id string) (err error) {
	if err := config.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
