package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURL string

func DBinstance() *mongo.Client {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Assign the value of MONGODB_URI
	mongoURL = os.Getenv("MONGODB_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		log.Fatal(err)
		fmt.Println("Database not connected successfully")
		return nil
	}

	fmt.Println("Database connected successfully")

	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("GoContacts").Collection(collectionName)

	return collection
}
