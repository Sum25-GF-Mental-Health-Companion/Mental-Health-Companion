package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"mental-health-api/database"
	"mental-health-api/handlers"
	"mental-health-api/middleware"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	database.Connect()

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)

	api.Use(middleware.JWTProtected())

	session := api.Group("/session")
	session.Get("/start", handlers.StartSession)
	session.Post("/end", handlers.EndSession)
	session.Post("/message", handlers.SendMessage)
	session.Get("/history", handlers.GetSessionHistory)

	log.Fatal(app.Listen(":3000"))
}
