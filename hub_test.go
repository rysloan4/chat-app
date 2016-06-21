package main

import (
	"chat/core"
	"chat/mocks"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockAuthenticator struct {
	mock.Mock
}

func (ma *MockAuthenticator) Authenticate(i interface{}) bool {
	return true
}

func TestLogUserOff(t *testing.T) {
	m := mocks.MockStorageManager{}
	ma := MockAuthenticator{}
	var hub = Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Conn),
		unregister:    make(chan *Conn),
		connections:   make(map[*Conn]bool),
		storageManger: &m,
		authenticator: &ma,
	}
	m.On("UpdateUserLastSeen", "foo", mock.Anything).Return(nil)
	hub.logUserOff("foo")
}

func TestSaveMessage(t *testing.T) {
	m := mocks.MockStorageManager{}
	ma := MockAuthenticator{}
	var hub = Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Conn),
		unregister:    make(chan *Conn),
		connections:   make(map[*Conn]bool),
		storageManger: &m,
		authenticator: &ma,
	}
	msg := &core.Message{
		UUID:      "foo",
		CreatedAt: time.Now(),
	}
	out, _ := json.Marshal(msg)
	m.On("InsertMessage", msg).Return(nil)
	hub.saveMessage(out)
}

func TestFetchUnreadMessages(t *testing.T) {
	m := mocks.MockStorageManager{}
	ma := MockAuthenticator{}
	var hub = Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Conn),
		unregister:    make(chan *Conn),
		connections:   make(map[*Conn]bool),
		storageManger: &m,
		authenticator: &ma,
	}
	m.On("GetUserByUsername", "foo").Return(nil)
	m.On("GetMessages", mock.Anything, "foo").Return(nil, []*core.Message{})
	hub.fetchUnreadMessages("foo")
}
