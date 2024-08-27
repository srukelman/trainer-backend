package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"trainer.seanrkelman.com/backend/routes"
)

func main() {
	router := gin.Default()
	router.GET("/activities", routes.GetActivities)
	err := godotenv.Load()
	mongoDBUser := os.Getenv("MONGO_DB_USER")
	mongoDBPassword := os.Getenv("MONGO_DB_PASSWORD")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://" + mongoDBUser + ":" + mongoDBPassword + "@cluster-trainer.fjyja.mongodb.net/?retryWrites=true&w=majority&appName=Cluster-trainer").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	router.Run("localhost:8080")
}
