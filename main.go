package main

import (
	"github.com/gin-gonic/gin"
	"trainer.seanrkelman.com/backend/routes"
	"trainer.seanrkelman.com/backend/server"
)

func main() {
	server.InitDb()
	router := gin.Default()
	router.GET("/activities", routes.GetActivities)
	router.POST("/activities", routes.CreateActivity)
	router.Run("localhost:8080")
}
