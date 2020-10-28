package routes

import (
	"github.com/IvanaKrizic/Dashboard/src/auth"
	"github.com/IvanaKrizic/Dashboard/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", auth.TokenAuth(), controllers.GetAllForthcomingEvents)
	router.GET("/home", auth.TokenAuth(), controllers.GetDailyEvents)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/logout", auth.TokenAuth(), controllers.Logout)

	groupUser := router.Group("/user", auth.TokenAuth())
	{
		groupUser.GET("/", controllers.GetUsers)
		groupUser.GET("/:id", controllers.GetUserByID)
		groupUser.PUT("/:id", controllers.UpdateUser)
		groupUser.DELETE("/:id", controllers.DeleteUser)
	}

	groupEvent := router.Group("/event", auth.TokenAuth())
	{
		groupEvent.GET("/", controllers.GetUserForthcomingEvents)
		groupEvent.POST("/", controllers.CreateEvents)
		groupEvent.GET("/:id", controllers.GetEventByID)
		groupEvent.PUT("/:id", controllers.UpdateEvent)
		groupEvent.DELETE("/:id", controllers.DeleteEvent)
	}

	return router
}
