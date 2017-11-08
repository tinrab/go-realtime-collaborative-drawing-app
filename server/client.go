package main

import "github.com/gorilla/websocket"

type Client struct {
	hub  *Hub
	conn *websocket.Conn
}

func newClient(hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		hub:  hub,
		conn: conn,
	}
}
