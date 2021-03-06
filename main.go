package main

import (
	"flag"
	"log"
	"net/http"

	"chat/authentication"
	"chat/data"
	"chat/handlers"
	"github.com/gorilla/mux"
	"os"
)

const (
	// LocalConnectionString is for connecting to mysql locally
	LocalConnectionString = "chat:chat@/chat?parseTime=true"

	// ProductionConnectionString is for connecting to mysql in production
	ProductionConnectionString = "b3fd3325d24b40:2761ce0f@tcp(us-cdbr-iron-east-04.cleardb.net:3306)/heroku_7dda9dbd4cbc075?parseTime=true"
)

// Get environmnet vars
var port = os.Getenv("PORT")
var env = os.Getenv("ENV")

// Set service address
var addr = flag.String("addr", ":"+port, "http service address")

// Global hub used by connections
var hub = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *Conn),
	unregister:  make(chan *Conn),
	connections: make(map[*Conn]bool),
}

func main() {
	var err error

	var connectionString string

	if env == "development" {
		connectionString = LocalConnectionString
	} else {
		connectionString = ProductionConnectionString
	}

	storageManager, err := data.NewMysqlStorageManager(connectionString)

	if err != nil {
		log.Fatalf("Could not initialize mysql storage manager: %s", err)
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
	chatHandler := handlers.NewChatHandler(storageManager)

	// Health
	r.HandleFunc("/health", chatHandler.IsHealthy).Methods(http.MethodGet)

	// Create User
	r.HandleFunc("/user", chatHandler.CreateUser).Methods(http.MethodPost)

	// Render chat page
	r.HandleFunc("/chat/{username}", chatHandler.ServeHome).Methods(http.MethodGet)

	// Render login page
	r.HandleFunc("/login", chatHandler.ServeLogin).Methods(http.MethodGet)

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

//TODO: auth tests, integration test, answer questions, submit
