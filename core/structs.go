package core

import (
	"time"
)

type (
	// Message is a struct to model a chat message
	Message struct {
		UUID         string    `json:"uuid"`
		FromUsername string    `json:"fromUsername"`
		ToUsername   string    `json:"toUsername"`
		Content      string    `json:"content"`
		CreatedAt    time.Time `json:"createdAt"`
	}

	// User is a struct to model a chat user
	User struct {
		UUID      string    `json:"uuid"`
		Username  string    `json:"username"`
		LastSeen  time.Time `json:"lastSeen"`
		CreatedAt time.Time `json:"createdAt"`
	}

	// CreateUserRequest is a struct to model a user creation request
	CreateUserRequest struct {
		Username string `json:"username"`
	}
)
