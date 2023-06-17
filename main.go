package main

import (
	"diary-api/controler"
	"diary-api/database"
	"diary-api/middleware"
	"diary-api/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controler.Register)
	publicRoutes.POST("/login", controler.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controler.AddEntry)
	protectedRoutes.GET("/entry", controler.GetAllEntries)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
