package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"mental-health-api/handlers"
	"mental-health-api/internal/llm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockLLM struct {
	Response string
	Err      error
}

func (m *mockLLM) SendMessage(_ context.Context, _ []llm.ChatMessage) (string, error) {
	return m.Response, m.Err
}

func TestNewMessageHandler_Success(t *testing.T) {
	mock := &mockLLM{Response: "Hello!"}
	handler := handlers.NewMessageHandler(mock)

	body := []byte(`{"text":"feel tired"}`)
	req := httptest.NewRequest(http.MethodPost, "/message", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp handlers.MessageResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}

	if resp.Response != "Hello!" {
		t.Errorf("unexpected response: %s", resp.Response)
	}
}

func TestNewMessageHandler_InvalidJSON(t *testing.T) {
	handler := handlers.NewMessageHandler(&mockLLM{})

	req := httptest.NewRequest(http.MethodPost, "/message", bytes.NewBuffer([]byte(`{invalid json}`)))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestNewMessageHandler_LLMError(t *testing.T) {
	mock := &mockLLM{Err: errors.New("API timeout")}
	handler := handlers.NewMessageHandler(mock)

	body := []byte(`{"text":"feel tired"}`)
	req := httptest.NewRequest(http.MethodPost, "/message", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}
