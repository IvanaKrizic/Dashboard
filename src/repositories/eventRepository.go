package repositories

import (
	"fmt"
	"time"

	"github.com/IvanaKrizic/Dashboard/src/config"
	"github.com/IvanaKrizic/Dashboard/src/models"
)

func GetAllForthcomingEvents(events *[]models.Event, startTime time.Time, endTime time.Time) (err error) {
	if err = config.DB.Where("date >= ? AND date <= ?", startTime, endTime).Find(events).Error; err != nil {
		return err
	}
	return nil
}

func GetDailyEvents(events *[]models.Event, date time.Time) (err error) {
	if err = config.DB.Where("date = ?", date).Find(events).Error; err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func GetUserForthcomingEvents(events *[]models.Event, id uint, startTime time.Time, endTime time.Time) (err error) {
	if err = config.DB.Where("user_id = ? AND date >= ? AND date <= ?", id, startTime, endTime).Find(events).Error; err != nil {
		return err
	}
	return nil
}

func CreateEvent(event *models.Event) (err error) {
	if err = config.DB.Create(event).Error; err != nil {
		return err
	}
	return nil
}

func GetEventByID(event *models.Event, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(event).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEvent(event *models.Event, id string) (err error) {
	config.DB.Select("EventType", "Note", "Date").Updates(event)
	return nil
}

func DeleteEvent(event *models.Event, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(event)
	return nil
}
