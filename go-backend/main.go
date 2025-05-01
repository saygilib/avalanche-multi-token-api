package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/saygilib/avalanche-project/go-backend/db"
	"github.com/saygilib/avalanche-project/go-backend/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file from go-backend/.env")
	}
	db.Init()
	r := gin.Default()
	handlers.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 API server started at http://localhost:" + port)
	r.Run(":" + port)
}
