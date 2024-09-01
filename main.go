package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"trainer.seanrkelman.com/backend/routes"
	"trainer.seanrkelman.com/backend/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	environment := os.Getenv("ENVIRONMENT")
	var reactClient string
	if environment == "development" {
		reactClient = os.Getenv("DEV_REACT_CLIENT")
	} else {
		reactClient = os.Getenv("PROD_REACT_CLIENT")
	}
	server.InitDb()
	router := gin.Default()
	trustedProxies := []string{reactClient}
	router.SetTrustedProxies(trustedProxies)
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{reactClient},
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
	router.GET("/workouts", routes.GetWorkouts)
	router.GET("/workouts/athlete/:athlete", routes.GetWorkoutsByAthlete)
	router.POST("/workouts", routes.CreateWorkout)
	router.Run("localhost:8080")
}
