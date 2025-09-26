package store

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnectionHelper() (mongo.Client, error) {
	godotenv.Load(".env") // load .env file if present
	uri := os.Getenv("mongodburi")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	// MongoDB connection logic here
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	// Handle errors and return the client
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Use the client...
	return *client, nil
}

func SaveEvent(message string) error {
	client, err := MongoConnectionHelper()
	if err != nil {
		return err
	}
	collection := client.Database("kafkadb").Collection("events")
	_, err = collection.InsertOne(context.TODO(), map[string]string{"message": message})
	if err != nil {
		log.Printf("Failed to insert document: %v", err)
		// Handle error
	}
	return err
}
