package llm_test

import (
	"context"
	"mental-health-api/internal/llm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendMessage_InvalidURL(t *testing.T) {
	_ = os.Setenv("PROXY_API_URL", "http://localhost:1234") // not running
	_ = os.Setenv("PROXY_API_KEY", "fake")

	client := llm.NewLLMClient()

	resp, err := client.SendMessage(context.Background(), []llm.ChatMessage{
		{Role: "user", Content: "test"},
	})

	require.Error(t, err)
	assert.Empty(t, resp)
}

func TestSendMessage_EmptyMessages(t *testing.T) {
	client := llm.NewLLMClient()

	resp, err := client.SendMessage(context.Background(), nil)

	require.Error(t, err)
	assert.Empty(t, resp)
}
