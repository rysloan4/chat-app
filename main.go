package main

import (
	"flag"
	"log"
	"net/http"

	"chat/authentication"
	"chat/data"
	"chat/handlers"
	"github.com/gorilla/mux"
)

var port = "8080"
var addr = flag.String("addr", ":"+port, "http service address")

func main() {
	var err error

	storageManager, err := data.NewMysqlStorageManager("chat:chat@/chat?parseTime=true")

	if err != nil {
		log.Fatal("Could not initialize mysql storage manager")
	}

	authenticator := authentication.NewUserNameAuthenticator(storageManager)

	handler := registerHandlers(storageManager)

	log.Printf("Service starting on port: %s", port)

	go hub.run(storageManager, authenticator)
	err = http.ListenAndServe(*addr, handler)

	if err != nil {
		log.Fatalf("Could not launch service: %v", err)
	}
}

// register handlers & setup routing
func registerHandlers(storageManager data.StorageManager) *mux.Router {
	r := mux.NewRouter()
	chatHanlder := handlers.NewChatHandler(storageManager)

	// Health
	r.HandleFunc("/health", chatHanlder.IsHealthy).Methods(http.MethodGet)

	// Create User
	r.HandleFunc("/user", chatHanlder.CreateUser).Methods(http.MethodPost)

	// Render chat page
	r.HandleFunc("/chat/{username}", chatHanlder.ServeHome).Methods(http.MethodGet)

	// Render login page
	r.HandleFunc("/login", chatHanlder.ServeLogin).Methods(http.MethodGet)

	// Websocket endpoint
	r.HandleFunc("/ws/{username}", handleWs)

	return r
}

// handles websocket requests from the peer
func handleWs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	conn := &Conn{send: make(chan []byte, 256), ws: ws, username: username}
	hub.register <- conn

	go conn.writePump()
	conn.readPump()
}
