package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMClient struct {
	client *resty.Client
	apiKey string
}

func NewLLMClient() *LLMClient {
	_ = godotenv.Load()
	fmt.Println("✅ PROXY_API_URL =", os.Getenv("PROXY_API_URL"))
	fmt.Println("✅ PROXY_API_KEY =", os.Getenv("PROXY_API_KEY"))
	client := resty.New().
		SetBaseURL(os.Getenv("PROXY_API_URL")). // https://api.proxyapi.ru/anthropic
		SetDebug(true)

	return &LLMClient{
		client: client,
		apiKey: os.Getenv("PROXY_API_KEY"), // Используем API-ключ от proxyapi.ru
	}
}

type AnthropicRequest struct {
	Model     string `json:"model"`
	MaxTokens int    `json:"max_tokens"`
	System    string `json:"system"`
	Messages  []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type AnthropicResponse struct {
	Content string `json:"content"` // proxyapi.ru возвращает упрощённую структуру
}

func (c *LLMClient) SendMessage(ctx context.Context, messages []ChatMessage) (string, error) {
	// Преобразуем в формат, который понимает proxyapi.ru
	var req AnthropicRequest
	req.Model = "claude-3-sonnet-20240229"
	req.MaxTokens = 3000
	for _, msg := range messages {
		if msg.Role == "system" {
			req.System = msg.Content // ✅ передаём system отдельно
			continue
		}

		req.Messages = append(req.Messages, struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	var result AnthropicResponse

	r, err := c.client.R().
		SetContext(ctx).
		SetHeader("Authorization", "Bearer "+c.apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&result).
		Post("/messages")

	if err != nil {
		return "", err
	}

	if r.IsError() {
		return "", fmt.Errorf("proxy error: %s", r.Status())
	}

	if result.Content == "" {
		return "", fmt.Errorf("empty response from Claude")
	}

	fmt.Println("[LLM RESPONSE]", result.Content)
	return result.Content, nil
}
