package sockets

import "fmt"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

// NewHub creates a new hub
func NewHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run handles registration and unregistration of clients
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("Registering client")
			h.clients[client] = true
		case client := <-h.unregister:
			delete(h.clients, client)
		}
	}
}
