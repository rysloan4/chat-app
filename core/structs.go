package core

import (
	"time"
)

type (
	Message struct {
		UUID	 	string		`json:"uuid"`
		FromUsername 	string 		`json:"fromUsername"`
		ToUsername 	string 		`json:"toUsername"`
		Content         string          `json:"content"`
		CreatedAt	time.Time	`json:"createdAt"`
	}

	User struct {
		UUID	 	string		`json:"uuid"`
		Username 	string 		`json:"username"`
		LastSeen	time.Time	`json:"lastSeen"`
		CreatedAt	time.Time	`json:"createdAt"`
	}

	CreateUserRequest struct {
		Username string                `json:"username"`
	}
)