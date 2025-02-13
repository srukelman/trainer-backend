package routes

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	cursor, err := server.GetMongoClient().Database("trainer").Collection("activities").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var activities []Activity
	if err := cursor.All(context.TODO(), &activities); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, activities)
}

func GetActivityByID(c *gin.Context) {
	id := c.Param("id")
	var activity Activity
	err := server.GetMongoClient().Database("trainer").Collection("activities").FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&activity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, activity)
}

func GetActivitiesByAthlete(c *gin.Context) {
	athlete := c.Param("athlete")
	cursor, err := server.GetMongoClient().Database("trainer").Collection("activities").Find(context.TODO(), bson.D{{Key: "athlete", Value: athlete}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var activities []Activity
	if err := cursor.All(context.TODO(), &activities); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, activities)
}

func GetMostRecentActivity(c *gin.Context) {
	athlete := c.Param("athlete")
	queryOptions := options.FindOne()
	queryOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	var activity Activity
	err := server.GetMongoClient().Database("trainer").Collection("activities").FindOne(context.TODO(), bson.D{{Key: "athlete", Value: athlete}}, queryOptions).Decode(&activity)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, activity)
}

func CreateActivity(c *gin.Context) {
	var activity Activity
	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var activityExists Activity
	_ = server.GetMongoClient().Database("trainer").Collection("activities").FindOne(context.TODO(), bson.D{{Key: "id", Value: activity.ID}}).Decode(&activityExists)
	if activityExists.ID != activity.ID {
		_, err := server.GetMongoClient().Database("trainer").Collection("activities").InsertOne(context.TODO(), activity)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(http.StatusCreated, activity)
	} else {
		c.IndentedJSON(http.StatusOK, activity)
	}

}

func UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var activity Activity
	if err := c.BindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.D{{Key: "id", Value: id}}
	if activity.Title != "" {
		update = append(update, bson.E{Key: "title", Value: activity.Title})
	}
	if activity.Athlete != "" {
		update = append(update, bson.E{Key: "athlete", Value: activity.Athlete})
	}
	if activity.Distance != 0 {
		update = append(update, bson.E{Key: "distance", Value: activity.Distance})
	}
	if activity.Time != 0 {
		update = append(update, bson.E{Key: "time", Value: activity.Time})
	}
	if !activity.Date.IsZero() {
		update = append(update, bson.E{Key: "date", Value: activity.Date})
	}

	if len(update) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := server.GetMongoClient().Database("trainer").Collection("activities").UpdateOne(context.TODO(), bson.D{{Key: "id", Value: id}}, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	_, err := server.GetMongoClient().Database("trainer").Collection("activities").DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
