package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	url := "ws://localhost:8080/ws" // Replace with the actual WebSocket server URL
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	message := "Hello, WebSocket Server!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("Write error:", err)
	}
	fmt.Println("Sent:", message)

	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Read error:", err)
	}
	fmt.Println("Received:", string(response))
}
