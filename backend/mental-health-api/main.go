package main

import (
	"log"
	"mental-health-api/database"
	"mental-health-api/handlers"
	"os"

	"mental-health-api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New())
	app.Use(logger.New())

	database.InitDatabase()
	handlers.SetDB(database.DB)

	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	app.Use(middleware.JWTProtected())

	session := app.Group("/session")
	session.Get("/start", handlers.StartSession)
	session.Post("/end", handlers.EndSession)

	app.Post("/message", handlers.SendMessage)
	app.Get("/sessions", handlers.GetSessionHistory)

	log.Fatal(app.Listen(":8080"))
}
