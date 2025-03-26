package main

import (
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan ClientMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan ClientMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			senderResponse, allResponse, err := handleMessage(message)
			if err != nil {
				message.sender.send <- fmt.Appendf(nil, "Error with message %v", err)
				return
			}
			message.sender.send <- []byte(senderResponse.Message)
			for client := range h.clients {
				select {
				case client.send <- []byte(allResponse.Message):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

type ClientMessage struct {
	sender  *Client
	content []byte
}

type ClientResponse struct {
	Message string `json:"message"`
}

func handleMessage(message ClientMessage) (ClientResponse, ClientResponse, error) {
	var senderResponse ClientResponse
	var allClientsResponse ClientResponse

	senderResponse = ClientResponse{
		Message: "you sent " + string(message.content),
	}

	allClientsResponse = ClientResponse{
		Message: string(message.content),
	}

	return senderResponse, allClientsResponse, nil
}
