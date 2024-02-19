package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sumitkoomar/JTWauthentication/database"
	"github.com/sumitkoomar/JTWauthentication/routes"
)

var PORT string

func main() {
	client := database.DBinstance()

	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	router := gin.New()

	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT = os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	router.Run(":" + PORT)

	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}

	fmt.Println("Disconnected from MongoDB.")

}
