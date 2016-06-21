package data

import (
	"chat/core"
	"time"
)

// StorageManager is an interface for storing data
type StorageManager interface {
	CleanUp()

	InsertMessage(m *core.Message) error

	GetMessages(from time.Time, userUUID string) ([]*core.Message, error)

	GetUserByUsername(username string) (*core.User, error)

	InsertUser(u *core.User) (*core.User, error)

	UpdateUserLastSeen(username string, lastSeen time.Time) error
}
