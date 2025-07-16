package handlers

import (
<<<<<<< HEAD
	"mental-health-api/models"
	"strconv"

	"time"
=======
	"mental-health-api/internal/llm"
>>>>>>> b00b55ae22b3329db506b9652771c063bab2b00b

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var llmClient = llm.NewLLMClient()

func Register(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and password are required",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": input.Email,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte("Mental Health"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not sign token",
		})
	}

	return c.JSON(fiber.Map{"token": signedToken})
}

func Login(c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 123,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	signedToken, err := token.SignedString([]byte("Mental Health"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not sign token",
		})
	}

	return c.JSON(fiber.Map{"token": signedToken})
}

func StartSession(c *fiber.Ctx) error {
	sessionID := uuid.New().String()
	return c.JSON(fiber.Map{"session_id": sessionID})
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
	var payload struct {
		Messages []models.Message `json:"messages"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	summary := models.Summary{
		FullSummary:       "Full summary",
		CompressedSummary: "Summary of " + strconv.Itoa(len(payload.Messages)) + " messages",
	}

	return c.JSON(fiber.Map{
		"full_summary":       summary.FullSummary,
		"compressed_summary": summary.CompressedSummary,
	})
}

func GetSessionHistory(c *fiber.Ctx) error {
	history := []models.Summary{
		{
			// SessionID:         "1234",
			// FullSummary:       "You talked about anxiety and routines.",
			CompressedSummary: "Talked about anxiety.",
		},
		{
			// SessionID:         "5678",
			// FullSummary:       "Session focused on productivity and habits.",
			CompressedSummary: "Habits & productivity.",
		},
	}

	return c.JSON(fiber.Map{"sessions": history})
}
