package handlers

import (
	"database/sql"
	"log"
	"mental-health-api/database/queries"
	"mental-health-api/internal/llm"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var llmClient = llm.NewLLMClient()

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func Register(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil || input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	q := queries.New(db)

	user, err := q.CreateUser(c.Context(), queries.CreateUserParams{
		Email:          input.Email,
		HashedPassword: input.Password,
		StudentStatus:  sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not create user"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	signed, err := token.SignedString([]byte("mental_health"))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "token error"})
	}

	return c.JSON(fiber.Map{"token": signed})
}

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil || input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	q := queries.New(db)

	user, err := q.GetUserByEmail(c.Context(), input.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "user not found"})
	}

	if user.HashedPassword != input.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "wrong password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	})

	signed, err := token.SignedString([]byte("mental_health"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "token error"})
	}

	return c.JSON(fiber.Map{"token": signed})
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
		{Role: "system", Content: "You are a caring psychologist. Help the student understand himself. Speak English."},
		{Role: "user", Content: body.Text},
	})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "LLM error: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"response": reply,
	})
}

func StartSession(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int32(claims["user_id"].(float64))

	q := queries.New(db)

	session, err := q.StartSession(c.Context(), sql.NullInt32{Int32: userID, Valid: true})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to start session",
		})
	}

	return c.JSON(session)
}

func EndSession(c *fiber.Ctx) error {
	var payload struct {
		SessionID int32             `json:"session_id"`
		Messages  []queries.Message `json:"messages"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	q := queries.New(db)

	err := q.EndSession(c.Context(), payload.SessionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to end session"})
	}

	full := "Full summary"
	compressed := "Summary of " + strconv.Itoa(len(payload.Messages)) + " messages"

	err = q.SaveSummary(c.Context(), queries.SaveSummaryParams{
		SessionID:         payload.SessionID,
		FullSummary:       sql.NullString{String: full, Valid: true},
		CompressedSummary: sql.NullString{String: compressed, Valid: true},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save summary"})
	}

	return c.JSON(fiber.Map{
		"full_summary":       full,
		"compressed_summary": compressed,
	})
}

func GetSessionHistory(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int32(claims["user_id"].(float64))

	q := queries.New(db)

	sessions, err := q.GetSessionsByUser(c.Context(), sql.NullInt32{Int32: userID, Valid: true})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch sessions"})
	}

	var summaries []fiber.Map
	for _, s := range sessions {
		summaries = append(summaries, fiber.Map{
			"session_id":         s.ID,
			"compressed_summary": "Session on " + s.StartedAt.Time.Format("2006-01-02 15:04"),
		})
	}

	return c.JSON(fiber.Map{"sessions": summaries})
}
