package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"chat/core"
	"chat/data"
	"github.com/gorilla/mux"
	"log"
)

type ChatHandler interface {
	IsHealthy(w http.ResponseWriter, r *http.Request)
	GetUserByUsername(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	ServeHome(w http.ResponseWriter, r *http.Request)
	ServeLogin(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	storageManager data.StorageManager
}

func NewChatHandler(storageManager data.StorageManager) ChatHandler {
	return &handler{storageManager}
}

// IsHealthy - handler to check if service is healthy
func (h *handler) IsHealthy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`OK`))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := core.CreateUserRequest{}

	if err := getStructFromRequest(r, &request); err != nil {
		log.Println(err)
	}

	user := core.User{
		Username: request.Username,
	}

	userResult, err := h.storageManager.InsertUser(&user)
	if err != nil {
		log.Println(err)
	}

	ret, err := h.marshalResponse(w, userResult)
	if err != nil {
		log.Println(err)
	}

	writeJSON(w, ret)
}

// GetUser - handler to get user
func (h *handler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, _ := h.storageManager.GetUserByUsername(username)

	ret, err := h.marshalResponse(w, user)
	if err != nil {
		return
	}
	writeJSON(w, ret)
}

func (h *handler) ServeHome(w http.ResponseWriter, r *http.Request) {
	var homeTemplate = template.Must(template.ParseFiles("templates/home.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func (h *handler) ServeLogin(w http.ResponseWriter, r *http.Request) {
	var loginTemplate = template.Must(template.ParseFiles("templates/login.html"))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	loginTemplate.Execute(w, r.Host)
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

func getStructFromRequest(r *http.Request, out interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(out)
}
