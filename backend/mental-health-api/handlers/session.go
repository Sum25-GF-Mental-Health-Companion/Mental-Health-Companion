package handlers

import (
	"mental-health-api/internal/llm"

	"github.com/gofiber/fiber/v2"
)

var llmClient = llm.NewLLMClient()

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
	type Request struct {
		Text string `json:"text"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	reply, err := llmClient.SendMessage(c.Context(), []llm.ChatMessage{
		{Role: "system", Content: "Ты — заботливый психолог. Помоги студенту разобраться в себе."},
		{Role: "user", Content: body.Text},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "LLM error: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": reply,
	})
}

func EndSession(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "session ended"})
}

func GetSessionHistory(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"sessions": []string{}})
}
