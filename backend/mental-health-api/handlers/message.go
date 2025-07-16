package handlers

import (
	"encoding/json"
	"mental-health-api/internal/llm"
	"net/http"
)

type MessageRequest struct {
	Text string `json:"text"`
}

type MessageResponse struct {
	Response string `json:"response"`
}

func NewMessageHandler(llmClient *llm.LLMClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req MessageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		reply, err := llmClient.SendMessage(r.Context(), []llm.ChatMessage{
			{Role: "system", Content: "You are a caring psychologist, help the student. Speak English."},
			{Role: "user", Content: req.Text},
		})
		if err != nil {
			http.Error(w, "LLM error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(MessageResponse{Response: reply})
	}
}
