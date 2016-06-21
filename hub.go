// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"chat/authentication"
	"chat/core"
	"chat/data"
	"encoding/json"
	"log"
	"time"
)

// Hub maintains the set of active connections and broadcasts messages to the
// connections.

type Hub struct {
	// Registered connections.
	connections map[*Conn]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *Conn

	// Unregister requests from connections.
	unregister chan *Conn

	// Storage manager used for saving data, authentication, etc...
	storageManger data.StorageManager

	// Authenticator used for authenticating a user
	authenticator authentication.Authenticator
}

var hub = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *Conn),
	unregister:  make(chan *Conn),
	connections: make(map[*Conn]bool),
}

func (h *Hub) run(storageManager data.StorageManager, authenticator authentication.Authenticator) {
	h.storageManger = storageManager
	h.authenticator = authenticator

	for {
		select {
		case conn := <-h.register:
			if h.authenticator.Authenticate(conn.username) {
				msgs := h.fetchUnreadMessages(conn.username)
				for _, msg := range msgs {
					out, err := json.Marshal(msg.Content)
					if err != nil {
						log.Println(err)
						break
					}
					conn.ws.WriteMessage(1, out)
				}
				h.connections[conn] = true
			} else {
				conn.ws.WriteMessage(1, []byte("Unauthenticated Username"))
				conn.ws.Close()
			}
		case conn := <-h.unregister:
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				close(conn.send)
				h.logUserOff(conn.username)
			}
		case message := <-h.broadcast:
			h.saveMessage(message)
			for conn := range h.connections {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(h.connections, conn)
					h.logUserOff(conn.username)
				}
			}
		}
	}
}

func (h *Hub) saveMessage(m []byte) {
	message := core.Message{}
	json.Unmarshal(m, &message)
	h.storageManger.InsertMessage(&message)
}

func (h *Hub) logUserOff(username string) {
	h.storageManger.UpdateUserLastSeen(username, time.Now())
}

func (h *Hub) fetchUnreadMessages(username string) []*core.Message {
	log.Println("fetching unread messages for: ", username)
	usr, _ := h.storageManger.GetUserByUsername(username)
	lastSeen := usr.LastSeen
	msgs, err := h.storageManger.GetMessages(lastSeen, username)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(msgs)
	return msgs
}
