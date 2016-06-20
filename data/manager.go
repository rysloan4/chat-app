package data

import (
	"chat/core"
	"time"
)

type StorageManager interface {
	CleanUp()

	InsertMessage(m *core.Message) error

	GetMessages(from time.Time, userUUID string) ([]*core.Message, error)

	GetUserByUsername(uuid string) (*core.User, error)

	InsertUser(u *core.User) (*core.User, error)

	UpdateUser(u *core.User) error
}
