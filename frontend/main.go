package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8080/ws" // Replace with the actual WebSocket server URL
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("websocket read error:", err)
				return
			}
			fmt.Println("\r", string(message))
			fmt.Print("> ")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			panic("fatal while reading input!")
		}
		input := scanner.Text()
		err = conn.WriteMessage(websocket.TextMessage, []byte(input))
		if err != nil {
			log.Fatal("Write error:", err)
		}
	}
}
