package controllers

import (
	"fmt"
	"net/http"

	"github.com/IvanaKrizic/Dashboard/src/auth"
	"github.com/IvanaKrizic/Dashboard/src/models"
	"github.com/IvanaKrizic/Dashboard/src/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := repositories.GetAllUsers(&users); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		fmt.Println(users)
		c.JSON(http.StatusOK, gin.H{"data": users})
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := repositories.GetUserByID(&user, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if err := repositories.GetUserByID(&user, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.BindJSON(&user)
	if err := repositories.UpdateUser(&user, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if err := repositories.DeleteUser(&user, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"User with id " + id: "is deleted."})
	}
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repositories.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

func Login(c *gin.Context) {
	var credentials models.UserCredentials
	var user models.User

	if err := c.ShouldBind(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong credentials"})
		return
	}

	if err := repositories.GetUserByEmail(&user, credentials.Email); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User with provided email doesn't exist"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	authData := models.AuthData{Token: token}
	if err := repositories.CreateAuth(&authData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": token})
}

func Logout(c *gin.Context) {
	tokenString := auth.ExtractToken(c.Request)
	repositories.DeleteAuth(tokenString)
	c.JSON(http.StatusOK, "Successfully logged out.")
}
