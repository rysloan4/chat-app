package handlers

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"chat/core"
	"time"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"errors"
	"bytes"
	"log"
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

func TestIsHealthy(t *testing.T) {
	handler := NewChatHandler(&MockStorageManager{})
	testServer := httptest.NewServer(http.HandlerFunc(handler.IsHealthy))
	defer testServer.Close()

	statusCode, _ := send(testServer.URL, http.MethodGet, "", nil)
	assert.Equal

}

// HTTP sending request and checking response code
func send(serviceURL string, method string, body string, headers map[string]string) (int, error) {
	req, err := http.NewRequest(method, serviceURL, bytes.NewBufferString(body))
	if err != nil {
		return 0, err
	}

	// setup header
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// retrieve response
	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		return response.StatusCode, errors.New(string(bodyByte))
	}

	return response.StatusCode, nil
}
