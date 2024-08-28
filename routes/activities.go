package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"trainer.seanrkelman.com/backend/server"

	"github.com/gin-gonic/gin"
)

type Activity struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Athlete  string    `json:"athlete"`
	Distance float64   `json:"distance"`
	Time     float64   `json:"time"`
	Date     time.Time `json:"date"`
}

func GetActivities(c *gin.Context) {
	cursor, error := server.GetMongoClient().Database("trainer").Collection("activities").Find(context.TODO(), bson.D{{}})
	if error != nil {
		log.Fatal(error)
	}
	var activities []Activity
	if err := cursor.All(context.TODO(), &activities); err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, activities)
}

func CreateActivity(c *gin.Context) {
	var activity Activity
	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activity.Date = time.Now()
	_, err := server.GetMongoClient().Database("trainer").Collection("activities").InsertOne(context.TODO(), activity)
	if err != nil {
		log.Fatal(err)
	}

	c.Status(http.StatusCreated)
}
