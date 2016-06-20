package main

import (
	"flag"
	"log"
	"net/http"

	"chat/data"
	"chat/handlers"
	"github.com/gorilla/mux"
)

var port = "8080"
var addr = flag.String("addr", ":"+port, "http service address")

func main() {
	var err error

	storageManager, err := data.NewMysqlStorageManager("chat:chat@/chat")

	if err != nil {
		log.Fatal("Could not initialize mysql storage manager")
	}

	handler := registerHandlers(storageManager)

	log.Printf("Service starting on port: %s", port)
	go hub.run()
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

	// User endpoints
	r.HandleFunc("/user/{username}", chatHanlder.GetUserByUsername).Methods(http.MethodGet)
	r.HandleFunc("/user", chatHanlder.CreateUser).Methods(http.MethodPost)

	// Render endpoints
	r.HandleFunc("/", chatHanlder.ServeHome).Methods(http.MethodGet)
	r.HandleFunc("/login", chatHanlder.ServeLogin).Methods(http.MethodGet)

	// Websocket endpoint
	r.HandleFunc("/ws/{uuid}", serveWs)

	return r
}
