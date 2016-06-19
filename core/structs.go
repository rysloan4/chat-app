package core

import "time"

type (
	Message struct {
		UUID	 	string		`json:"uuid"`
		FromUUID 	string 		`json:"fromUUID"`
		ToUUID 		string 		`json:"toUUID"`
		IsRead		bool		`json:"isRead"`
		CreatedAt	time.Time	`json:"createdAt"`
		ReadAt		time.Time	`json:"readAt"`
	}

	User struct {
		UUID	 	string		`json:"uuid"`
		Username 	string 		`json:"username"`
		LastSeen	time.Time	`json:"lastSeen"`
		CreatedAt	time.Time	`json:"createdAt"`
	}
)