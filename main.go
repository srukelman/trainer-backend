package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"trainer.seanrkelman.com/backend/routes"
	"trainer.seanrkelman.com/backend/server"
)

func main() {
	server.InitDb()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))
	router.GET("/activities", routes.GetActivities)
	router.GET("/activities/:id", routes.GetActivityByID)
	router.GET("/activities/athlete/:athlete", routes.GetActivitiesByAthlete)
	router.GET("/activities/most-recent/:athlete", routes.GetMostRecentActivity)
	router.POST("/activities", routes.CreateActivity)
	router.PUT("/activities/:id", routes.UpdateActivity)
	router.DELETE("/activities/:id", routes.DeleteActivity)
	router.Run("localhost:8080")
}
