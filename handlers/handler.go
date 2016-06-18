package handlers

import (
	"net/http"
	"chat/data"
)

// BusinessHandler - interface for handlers for the business service
type ChatHandler interface {
	IsHealthy(w http.ResponseWriter, r *http.Request)
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