package main

import (
	"log"

<<<<<<< HEAD
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

=======
>>>>>>> b00b55ae22b3329db506b9652771c063bab2b00b
	"mental-health-api/database"
	_ "mental-health-api/database"
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
