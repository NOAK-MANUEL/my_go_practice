package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		println("http request: ", r)
		return true
	},
}

type Hub struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

func newHub() *Hub {
	return &Hub{clients: make(map[*websocket.Conn]bool)}
}

func (h *Hub) addClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[conn] = true
	println("Added new client: ", &conn)
}

func (h *Hub) removeClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, conn)

	println("Client removed: ", &conn)
}

func (h *Hub) broadCast(sender *websocket.Conn, msgType int, msg []byte) {

	h.mu.Lock()

	defer h.mu.Unlock()

	for client := range h.clients {
		if sender == client {
			continue
		}

		if err := client.WriteMessage(msgType, msg); err != nil {
			println("write error:", err)
		}
	}
}

func main() {
	h := newHub()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		h.addClient(conn)
		defer conn.Close()
		defer h.removeClient(conn)

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				println("Error occurred at 77: ", err)
				break
			}

			println("Relaying message: ", string(msg))

			h.broadCast(conn, msgType, msg)

		}

	})
	println("Listening at /ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
