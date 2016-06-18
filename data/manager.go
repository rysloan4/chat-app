package data

import (
	"chat/core"
	"time"
)

type StorageManager interface {
	CleanUp()

	PutMessage(m *core.Message) error

	GetMessages(from time.Time, until time.Time, userUUID string) ([]*core.Message, error)

	GetUser(uuid string) (*core.User, error)

	PutUser(u *core.User) error

	PostUser(u *core.User) error
}
