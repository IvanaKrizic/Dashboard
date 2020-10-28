package controllers

import (
	"net/http"
	"time"

	"github.com/IvanaKrizic/Dashboard/src/models"
	"github.com/IvanaKrizic/Dashboard/src/repositories"
	"github.com/gin-gonic/gin"
)

const (
	dateLayout = "21/10/2020"
)

func GetAllForthcomingEvents(c *gin.Context) {
	var events []models.Event
	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endTime := startTime.AddDate(0, 0, 7)

	if err := repositories.GetAllForthcomingEvents(&events, startTime, endTime); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": events})
	}
}

func GetDailyEvents(c *gin.Context) {
	var events []models.Event
	var t time.Time

	input := c.Request.URL.Query()
	if len(input["date"]) > 0 {
		t, _ = time.Parse(dateLayout, input["date"][0])
	} else {
		t = time.Now()
	}
	date := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	if err := repositories.GetDailyEvents(&events, date); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": events})
}

func GetUserForthcomingEvents(c *gin.Context) {
	var events []models.Event
	t := time.Now()
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	endTime := startTime.AddDate(0, 0, 7)
	userID := uint(c.MustGet("userId").(float64))

	if err := repositories.GetUserForthcomingEvents(&events, userID, startTime, endTime); err != nil {
		c.JSON(http.StatusNotFound, events)
	} else {
		c.JSON(http.StatusOK, events)
	}
}

func CreateEvents(c *gin.Context) {
	var eventDTO models.EventDTO

	if err := c.ShouldBindJSON(&eventDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := uint(c.MustGet("userId").(float64))
	start, _ := time.Parse(time.RFC3339, eventDTO.StartDate)
	end, _ := time.Parse(time.RFC3339, eventDTO.EndDate)

	for date := start; date.After(end) == false; date = date.AddDate(0, 0, 1) {
		event := models.Event{EventType: eventDTO.EventType, Note: eventDTO.Note, UserID: userID, Date: date}
		if err := repositories.CreateEvent(&event); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, "Event(s) successfully created.")
}

func GetEventByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var event models.Event
	if err := repositories.GetEventByID(&event, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, event)
	}
}

func UpdateEvent(c *gin.Context) {
	var event models.Event
	id := c.Params.ByName("id")
	if err := repositories.GetEventByID(&event, id); err != nil {
		c.JSON(http.StatusNotFound, event)
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repositories.UpdateEvent(&event, id); err != nil {
		c.JSON(http.StatusNotFound, event)
	} else {
		c.JSON(http.StatusOK, event)
	}
}

func DeleteEvent(c *gin.Context) {
	var event models.Event
	id := c.Params.ByName("id")
	if err := repositories.DeleteEvent(&event, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, "Event with id "+id+" is deleted")
	}
}
