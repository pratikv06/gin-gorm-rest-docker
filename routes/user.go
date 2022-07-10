package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pratikv06/gin-gorm-rest-docker/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}
	