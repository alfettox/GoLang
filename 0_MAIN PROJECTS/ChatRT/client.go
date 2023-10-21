/*
Author: Giovanni De Franceschi
*/
package main

import (
    "fmt"
    "os"
    "bufio"
    "github.com/gorilla/websocket"
)

func main() {
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080", nil)
    if err != nil {
        fmt.Println("Error connecting to the server:", err)
        os.Exit(1)
    }
    defer conn.Close()

    fmt.Print("Enter your username: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    username := scanner.Text()

    err = conn.WriteMessage(websocket.TextMessage, []byte(username))
    if err != nil {
        fmt.Println("Error sending username:", err)
        os.Exit(1)
    }

    go receiveMessages(conn)

    for {
        fmt.Print("Type a message: ")
        scanner.Scan()
        message := scanner.Text()
        err = conn.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            fmt.Println("Error sending message:", err)
            os.Exit(1)
        }
    }
}

func receiveMessages(conn *websocket.Conn) {
    for {
        _, p, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error receiving message:", err)
            return
        }
        message := string(p)
        fmt.Println(message)
    }
}
