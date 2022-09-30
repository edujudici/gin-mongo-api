package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser())
	router.GET("/users/:userId", controllers.GetAUser())
	router.PUT("/users/:userId", controllers.EditAUser())
	router.DELETE("/users/:userId", controllers.DeleteAUser())
	router.GET("/users", controllers.GetAllUsers())
}
