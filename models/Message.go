package models

// Structure of the Messages
type Message struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	Message  string `json:"message"`
}
