package handlers

import (
	"bytes"
	"chat/core"
	"chat/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsHealthy(t *testing.T) {
	handler := NewChatHandler(&mocks.MockStorageManager{})
	testServer := httptest.NewServer(http.HandlerFunc(handler.IsHealthy))
	defer testServer.Close()

	req, _ := http.NewRequest(http.MethodGet, testServer.URL, bytes.NewBufferString(""))
	client := &http.Client{}
	response, _ := client.Do(req)
	assert.Equal(t, response.StatusCode, 200)

}

func TestCreateUser(t *testing.T) {
	m := &mocks.MockStorageManager{}
	handler := NewChatHandler(m)
	testServer := httptest.NewServer(http.HandlerFunc(handler.CreateUser))
	defer testServer.Close()
	m.On("InsertUser", &core.User{}).Return(nil)

	req, _ := http.NewRequest(http.MethodGet, testServer.URL, bytes.NewBufferString(""))
	client := &http.Client{}
	response, _ := client.Do(req)
	log.Println(response)
	assert.Equal(t, response.StatusCode, 200)
}
