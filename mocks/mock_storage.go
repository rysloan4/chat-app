package mocks

import (
	"github.com/stretchr/testify/mock"
	"chat/core"
	"time"
)

// MockStorage for tests
type MockStorageManager struct {
	mock.Mock
}

func (msm *MockStorageManager) CleanUp() {}

func (msm *MockStorageManager) InsertMessage(m *core.Message) error {
	ret := msm.Called(m)

	r0 := ret.Error(0)

	return r0
}

func (msm *MockStorageManager) GetMessages(from time.Time, userUUID string) ([]*core.Message, error) {
	ret := msm.Called(from, userUUID)

	r0 := ret.Error(0)

	return []*core.Message{}, r0
}


func (msm *MockStorageManager) GetUserByUsername(username string) (*core.User, error) {
	ret := msm.Called(username)

	r0 := ret.Error(0)

	return &core.User{}, r0
}

func (msm *MockStorageManager) InsertUser(u *core.User) (*core.User, error) {
	ret := msm.Called(u)

	r0 := ret.Error(0)

	return &core.User{}, r0
}

func (msm *MockStorageManager) UpdateUserLastSeen(username string, lastSeen time.Time) error {
	ret := msm.Called(username, lastSeen)

	r0 := ret.Error(0)

	return r0
}

