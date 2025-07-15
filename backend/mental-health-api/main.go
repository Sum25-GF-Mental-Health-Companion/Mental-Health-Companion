package main

import (
	"log"

	"mental-health-api/database"
	"mental-health-api/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ .env file not loaded")
	}

	log.Println("✅ PROXY_API_URL =", os.Getenv("PROXY_API_URL"))
	log.Println("✅ PROXY_API_KEY =", os.Getenv("PROXY_API_KEY"))
	app := fiber.New()
	app.Use(logger.New())

	database.Connect()

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	// api.Use(middleware.JWTProtected())

	session := api.Group("/session")
	session.Get("/start", handlers.StartSession)
	session.Post("/end", handlers.EndSession)
	session.Post("/message", handlers.SendMessage)
	session.Get("/history", handlers.GetSessionHistory)

	log.Fatal(app.Listen(":3000"))
}
