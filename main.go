package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pratikv06/gin-gorm-rest-docker/config"
	"github.com/pratikv06/gin-gorm-rest-docker/routes"
)

func main() {
	// Creating a gin router
	router := gin.New()

	// Connecting to postgres database
	config.Connect()

	// adding routes
	routes.UserRoute(router)

	// Running the server
	router.Run(":8080")
}
