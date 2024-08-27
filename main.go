package main

import (
	"github.com/gin-gonic/gin"
	"trainer.seanrkelman.com/backend/routes"
)

func main() {
	router := gin.Default()
	router.GET("/activities", routes.GetActivities)
	router.Run("localhost:8080")
}
