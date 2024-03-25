package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/proGabby/4genz/domain/entity"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketRoutes(r *mux.Router, feedChan *chan *entity.Feed) {

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		defer conn.Close()

		go func(conn *websocket.Conn) {

			defer conn.Close()

			for {
				// get created feed from the feed channel
				if feedChan != nil {
					feed := <-*feedChan
					err := conn.WriteJSON(feed)
					if err != nil {
						log.Println("Error sending feed to client:", err)
						return
					}
				}
			}
		}(conn)

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message from client:", err)
				return
			}
			// Handle client messages here if needed
			log.Printf("Received message from client: %s\n", p)

			// Echo back the received message (optional)
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println("Error echoing message back to client:", err)
				return
			}
		}
	})

}
