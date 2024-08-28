package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoDBUser := os.Getenv("MONGO_DB_USER")
	mongoDBPassword := os.Getenv("MONGO_DB_PASSWORD")
	uri := "mongodb+srv://" + mongoDBUser + ":" + mongoDBPassword + "@cluster-trainer.fjyja.mongodb.net/?retryWrites=true&w=majority&appName=Cluster-trainer"
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	mongoClient = client
	fmt.Println("Connected to MongoDB!")
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}
