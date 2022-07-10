package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pratikv06/gin-gorm-rest-docker/config"
	"github.com/pratikv06/gin-gorm-rest-docker/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
	// c.String(200, "Hello Gin!")
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(201, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Where("id = ?", id).Delete(&user)
	c.JSON(204, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	config.DB.Where("id = ?", id).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(201, &user)
}
