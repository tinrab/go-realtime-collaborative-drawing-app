package main

import (
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc("/ws", hub.handle)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
