package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"chat/data"
)

// BusinessHandler - interface for handlers for the business service
type ChatHandler interface {
	IsHealthy(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetMessages(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

// Handler - implements the handlers for the business http service.
type handler struct{
	storageManager data.StorageManager
}

// NewBusinessHandler - initializes new business handler
func NewChatHandler(storageManager data.StorageManager) ChatHandler {
	return &handler{storageManager}
}

// IsHealthy - handler to check if service is healthy
func (h *handler) IsHealthy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`OK`))
}

// GetUser - handler to get user
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	user, _ := h.storageManager.GetUser(uuid)

	ret, err := h.marshalResponse(w, user)
	if err != nil {
		return
	}
	writeJSON(w, ret)
}

func (h *handler) marshalResponse(w http.ResponseWriter, response interface{}) ([]byte, error) {
	ret, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func writeJSON(w http.ResponseWriter, resp []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}