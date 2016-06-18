package data

import (
	"chat/core"
	"time"
)

type MysqlStorageManager struct {}

func NewMysqlStorageManager() (StorageManager, error) {
	return &MysqlStorageManager{}, nil
}

func (msm *MysqlStorageManager) CleanUp() {}

func (msm *MysqlStorageManager) PutMessage(m *core.Message) error {
	return nil
}

func (msm *MysqlStorageManager) GetMessages(from time.Time, until time.Time, userUUID string) ([]*core.Message, error) {
	return nil, nil
}

func (msm *MysqlStorageManager) GetUser(uuid string) (*core.User, error) {
	return nil, nil
}

func (msm *MysqlStorageManager) PutUser(u *core.User) error {
	return nil
}

func (msm *MysqlStorageManager) PostUser(u *core.User) error {
	return nil
}