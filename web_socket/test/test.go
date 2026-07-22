package test

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

func newHub() *Hub {
	return &Hub{clients: make(map[*websocket.Conn]bool)}
}
func (h *Hub) addClient(conn *websocket.Conn) {
	h.mu.Lock()         // wait here until no one else holds the lock, then take it
	defer h.mu.Unlock() // release it automatically when this function returns
	h.clients[conn] = true
	log.Printf("client connected, total: %d", len(h.clients))
}

func (h *Hub) removeClient(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, conn)
	log.Printf("client disconnected, total: %d", len(h.clients))

}

//broadcast sends msg to every client except the sender

func (h *Hub) broadcast(sender *websocket.Conn, msgType int, msg []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for conn := range h.clients {
		if conn == sender {
			continue
		}
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println("write error:", err)
		}
	}
}

func test() {
	hub := newHub()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		println("conn: ", conn)
		if err != nil {
			log.Println("upgrade error:", err)
			return
		}

		defer conn.Close()

		hub.addClient(conn)
		defer hub.removeClient(conn)

		for {
			msgType, msg, err := conn.ReadMessage() // BLOCKS here, doing nothing, until a message arrives

			if err != nil {
				println("read error (client likely disconnected)")
				break
			}

			log.Printf("relaying message: %s", msg) // only prints when a message actually arrives
			hub.broadcast(conn, msgType, msg)
		}

	})
	log.Println("Signaling server listening on :8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
