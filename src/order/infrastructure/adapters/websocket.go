package adapters

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketRepository struct {
	clients   map[*websocket.Conn]bool
	broadcast chan string
	upgrader  websocket.Upgrader
}

func NewWebSocketRepository() *WebSocketRepository {
	return &WebSocketRepository{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan string),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (ws *WebSocketRepository) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ws.clients[conn] = true
	log.Println("Nuevo cliente conectado")

	for {
		var msg string
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(ws.clients, conn)
			break
		}
		ws.BroadcastMessage(msg)
	}
}

func (ws *WebSocketRepository) HandleMessages() {
	for {
		msg := <-ws.broadcast
		for client := range ws.clients {
			err := client.WriteJSON(fmt.Sprintf("Estado de la orden: %s", msg))
			if err != nil {
				log.Printf("Error enviando: %v", err)
				client.Close()
				delete(ws.clients, client)
			}
		}
	}
}

func (ws *WebSocketRepository) BroadcastMessage(message string) {
	ws.broadcast <- message
}
