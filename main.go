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
	router.GET("/activities/:id", routes.GetActivityByID)
	router.PUT("/activities/:id", routes.UpdateActivity)
	router.DELETE("/activities/:id", routes.DeleteActivity)
	router.Run("localhost:8080")
}
