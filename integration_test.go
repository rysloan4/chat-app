package main

import (
	"chat/data"
	"testing"
	//"chat/handlers"
	"bytes"
	"chat/authentication"
	"chat/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"time"
)

type DoNothingResponseWriter struct {
}

func (d DoNothingResponseWriter) Header() http.Header {
	return http.Header{}
}

func (d DoNothingResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (d DoNothingResponseWriter) WriteHeader(int) {

}

var storageManager, _ = data.NewMysqlStorageManager("chat:chat@/chat?parseTime=true")
var auth = authentication.NewUserNameAuthenticator(storageManager)

func TestCanCreateUser(t *testing.T) {
	var uuid, username string
	var lastSeen, createdAt time.Time
	handler := handlers.NewChatHandler(storageManager)
	w := DoNothingResponseWriter{}
	var jsonStr = []byte(`{"username":"test_user"}`)
	req, _ := http.NewRequest(http.MethodPost, "localhost:8080/user", bytes.NewBuffer(jsonStr))
	handler.CreateUser(w, req)
	db := storageManager.GetDB()
	row := db.QueryRow("SELECT * FROM USER WHERE username = 'test_user';")
	row.Scan(&uuid, &username, &lastSeen, &createdAt)
	assert.Equal(t, "test_user", username)
	db.Query("DELETE FROM USER WHERE username = 'test_user';")
}

func TestCanRegisterConnection(t *testing.T) {
	go hub.run(storageManager, auth)
	w := DoNothingResponseWriter{}
	r, _ := http.NewRequest(http.MethodGet, "ws://localhost:8080/ws/rsloan", bytes.NewBufferString(""))
	ws, _ := upgrader.Upgrade(w, r, nil)
	conn := &Conn{send: make(chan []byte, 256), ws: ws, username: "rsloan"}
	hub.register <- conn
	//TODO fix so connection actually registers
}

func TestCanRejectUnauthUser(t *testing.T) {
	//TODO implement
}

func TestCanSendMessage(t *testing.T) {
	//TODO implement
}
