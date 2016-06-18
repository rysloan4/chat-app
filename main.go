package main

import (
	"flag"
	"net/http"

	"log"
	"chat/handlers"
	"github.com/gorilla/mux"
	"chat/data"
)

var port = "8080"
var addr = flag.String("addr", ":" + port, "http service address")

func main() {
	var err error

	storageManager, err := data.NewMysqlStorageManager()

	if err != nil {
		log.Fatal("Could not initialize mysql storage manager")
	}

	handler := registerHandlers(storageManager)

	log.Printf("Service starting on port: %s", port)

	err = http.ListenAndServe(*addr, handler)

	if err != nil {
		log.Fatalf("Could not launch service: %v", err)
	}
}

// register handlers & setup routing
func registerHandlers(storageManager data.StorageManager) *mux.Router {
	r := mux.NewRouter()
	chatHanlder := handlers.NewChatHandler(storageManager)
	r.HandleFunc("/health", chatHanlder.IsHealthy).Methods(http.MethodGet)
	return r
}
