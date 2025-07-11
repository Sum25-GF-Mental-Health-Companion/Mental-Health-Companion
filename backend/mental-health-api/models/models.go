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
	SessionID string
	Sender    string
	Content   string
}

type Summary struct {
	SessionID         string
	FullSummary       string
	CompressedSummary string
}
