/*
Author: Giovanni De Franceschi
*/
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

var broadcast = make(chan string)

func main() {
	http.HandleFunc("/", handleConnection)
	go handleMessages()

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	clients[conn] = true

	for {
		_, p, err := conn.ReadMessage() 
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			delete(clients, conn)
			break
		}

		message := string(p)
		fmt.Println("Received message:", message)
		broadcast <- message
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		fmt.Println("Broadcasting message:", message)
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				fmt.Println("Error broadcasting message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
