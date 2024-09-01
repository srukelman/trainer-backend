package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"trainer.seanrkelman.com/backend/server"
)

type Interval struct {
	Distance float64 `json:"distance"`
	Time     float64 `json:"time"`
	RestTime float64 `json:"restTime"`
	Pace     float64 `json:"pace"`
	Type     string  `json:"type"`
}

type Workout struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Athlete   string     `json:"athlete"`
	Distance  float64    `json:"distance"`
	Time      float64    `json:"time"`
	Date      time.Time  `json:"date"`
	Type      string     `json:"type"`
	Intervals []Interval `json:"intervals"`
}

func GetWorkouts(c *gin.Context) {
	cursor, err := server.GetMongoClient().Database("trainer").Collection("workouts").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var workouts []Workout
	if err := cursor.All(context.TODO(), &workouts); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, workouts)
}

func GetWorkoutsByAthlete(c *gin.Context) {
	athlete := c.Param("athlete")
	cursor, err := server.GetMongoClient().Database("trainer").Collection("workouts").Find(context.TODO(), bson.D{{Key: "athlete", Value: athlete}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var workouts []Workout
	if err := cursor.All(context.TODO(), &workouts); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, workouts)
}

func CreateWorkout(c *gin.Context) {
	var workout Workout
	if err := c.BindJSON(&workout); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := server.GetMongoClient().Database("trainer").Collection("workouts").InsertOne(context.TODO(), workout)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, workout)
}
