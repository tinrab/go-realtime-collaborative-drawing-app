package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				hub.onDisconnect(client)
			}
		}
	}
}

func (hub Hub) send(data []byte, ignore *Client) {
	for client := range hub.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}

func (hub *Hub) onConnect(client *Client) {
	hub.clients[client] = true
	log.Println("client connected: ", client.socket.RemoteAddr())
}

func (hub *Hub) onDisconnect(client *Client) {
	client.close()
	delete(hub.clients, client)

	log.Println("client disconnected: ", client.socket.RemoteAddr())
}

func (hub *Hub) onMessage(data []byte, client *Client) {
	log.Println("onMessage: ", string(data))
	hub.send(data, client)
}

func (hub *Hub) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "could not upgrade", http.StatusInternalServerError)
		return
	}
	client := newClient(hub, socket)
	hub.register <- client

	go client.read()
	go client.write()
}
