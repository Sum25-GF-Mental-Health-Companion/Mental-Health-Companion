package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"mental-health-api/database"
	_ "mental-health-api/database"
	"mental-health-api/handlers"
	"mental-health-api/middleware"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	database.InitDatabase()

	auth := app.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	app.Use(middleware.JWTProtected())

	session := app.Group("/session")
	session.Get("/start", handlers.StartSession)
	session.Post("/end", handlers.EndSession)
	app.Post("/message", handlers.SendMessage)
	// session.Get("/history", handlers.GetSessionHistory)
	app.Get("/sessions", handlers.GetSessionHistory)

	log.Fatal(app.Listen(":8080"))
}
