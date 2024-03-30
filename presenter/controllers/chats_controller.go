package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/proGabby/4genz/domain/entity"
	chatusecase "github.com/proGabby/4genz/domain/usecase/chat_usecase"
	"github.com/proGabby/4genz/utils"
)

type ChatController struct {
	chatUsecases chatusecase.ChatUsecases
	upgrader     websocket.Upgrader
	clients      map[*websocket.Conn]int
	clientsMu    sync.Mutex
}

func NewChatController() *ChatController {
	return &ChatController{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		clients:   make(map[*websocket.Conn]int),
		clientsMu: sync.Mutex{},
	}
}

func (c *ChatController) ManageChat(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value("user").(*entity.User)

	if user == nil || !ok {
		utils.HandleError(map[string]interface{}{
			"error": "user not authenticated",
		}, http.StatusBadRequest, w)
		return
	}

	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	c.clientsMu.Lock()
	c.clients[conn] = user.Id
	c.clientsMu.Unlock()

	for {
		// Read message from client
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from client:", err)
			break
		}

		var msg entity.Message
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		// Find recipient connection
		var recipientConn *websocket.Conn
		c.clientsMu.Lock()
		for conn, id := range c.clients {
			if id == msg.ReceiverId {
				recipientConn = conn
				break
			}
		}
		c.clientsMu.Unlock()

		// If recipient connection found, send message to recipient
		if recipientConn != nil {
			msgBytes, err := json.Marshal(msg)
			if err != nil {
				log.Println("Error marshalling message:", err)
				continue
			}
			if err := recipientConn.WriteMessage(messageType, msgBytes); err != nil {
				log.Println("Error sending message to recipient:", err)
			}
		} else {
			log.Println("Recipient not found")
		}

		//save message to database x
		err = c.chatUsecases.SaveChatMsg.Execute(&msg)
		if err != nil {
			log.Println("Error saving message to database:", err)
		}
	}

	c.clientsMu.Lock()
	delete(c.clients, conn)
	c.clientsMu.Unlock()

}
