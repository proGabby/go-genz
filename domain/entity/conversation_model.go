package entity

import "time"

type Conversation struct {
	Id         int    `json:"id"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	CreatedAt  string `json:"created_at"`
}

func NewConversation(id, senderId, receiverId int) *Conversation {
	return &Conversation{
		Id:         id,
		SenderId:   senderId,
		ReceiverId: receiverId,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
}
