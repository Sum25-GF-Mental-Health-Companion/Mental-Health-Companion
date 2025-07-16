package handlers

import (
	"database/sql"
	"log"
	"mental-health-api/database/queries"
	"mental-health-api/internal/llm"
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
		{Role: "system", Content: "Ты — заботливый психолог. Помоги студенту разобраться в себе."},
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

// Updated helper function (assuming messages only have Content)
func joinMessages(messages []queries.Message) string {
	var result string
	for _, m := range messages {
		result += m.Content + "\n" // Just concatenate content
	}
	return result
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

	conversation := joinMessages(payload.Messages)

	// Generate summaries using the LLM
	fullSummary, err := llmClient.SendMessage(c.Context(), []llm.ChatMessage{
		{Role: "system", Content: "Summarize this therapy session in detail in Russian."},
		{Role: "user", Content: joinMessages(payload.Messages)}, // Combine all messages into a single string
	})

	if err != nil {
		log.Println("LLM summary error:", err)
		fullSummary = "Не удалось создать резюме сессии"
	}

	compressedSummary, err := llmClient.SendMessage(c.Context(), []llm.ChatMessage{
		{Role: "system", Content: "Ты — психолог. Создай очень краткое (1 предложение) резюме этой сессии на русском."},
		{Role: "user", Content: conversation},
	})
	if err != nil {
		log.Println("LLM compressed summary error:", err)
		compressedSummary = "Краткое резюме недоступно"
	}

	// Save to database
	err = q.SaveSummary(c.Context(), queries.SaveSummaryParams{
		SessionID:         payload.SessionID,
		FullSummary:       sql.NullString{String: fullSummary, Valid: true},
		CompressedSummary: sql.NullString{String: compressedSummary, Valid: true},
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save summary"})
	}

	return c.JSON(fiber.Map{
		"full_summary":       fullSummary,
		"compressed_summary": compressedSummary,
	})
}

func GetSessionHistory(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int32(claims["user_id"].(float64))

	q := queries.New(db)

	// 1. Get all sessions for the user
	sessions, err := q.GetSessionsByUser(c.Context(), sql.NullInt32{Int32: userID, Valid: true})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch sessions"})
	}

	// 2. For each session, get its summary (if exists)
	var sessionSummaries []fiber.Map
	for _, session := range sessions {
		// Try to get summary - this assumes your summaries table has a session_id foreign key
		summary, err := q.GetSummaryBySession(c.Context(), session.ID)
		if err != nil {
			// If no summary exists, create a basic one
			log.Println("basic summary")
			sessionSummaries = append(sessionSummaries, fiber.Map{
				"session_id":         session.ID,
				"started_at":         session.StartedAt.Time.Format("2006-01-02 15:04"),
				"compressed_summary": "Сессия от " + session.StartedAt.Time.Format("2006-01-02"),
				"full_summary":       "Детали сессии недоступны",
			})
			continue
		}

		// If summary exists, use it
		sessionSummaries = append(sessionSummaries, fiber.Map{
			"session_id":         session.ID,
			"started_at":         session.StartedAt.Time.Format("2006-01-02 15:04"),
			"compressed_summary": summary.CompressedSummary.String,
			"full_summary":       summary.FullSummary.String,
		})
	}

	log.Println(len(sessionSummaries))

	return c.JSON(fiber.Map{"sessions": sessionSummaries})
}
