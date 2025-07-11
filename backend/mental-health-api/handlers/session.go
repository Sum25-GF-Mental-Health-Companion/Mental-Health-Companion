package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "registered"})
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"token": ""})
}

func StartSession(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"session_id": "uuid"})
}

func SendMessage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"reply": "AI response"})
}

func EndSession(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "session ended"})
}

func GetSessionHistory(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"sessions": []string{}})
}
