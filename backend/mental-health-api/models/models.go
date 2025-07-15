package models

type User struct {
	ID       int
	Email    string
	Password string
}

type Session struct {
	ID     string
	UserID int
}

type Message struct {
	// SessionID string `json:"session_id,omitempty"`
	// Sender  string `json:"sender"`
	// Content string `json:"content"`
	Content string `json:"message"`
}

type Summary struct {
	// SessionID         string `json:"session_id,omitempty"`
	FullSummary       string `json:"full_summary"`
	CompressedSummary string `json:"compressed_summary"`
}
