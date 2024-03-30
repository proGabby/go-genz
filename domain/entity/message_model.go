package entity

type Message struct {
	Id             int    `json:"id"`
	ConversationId int    `json:"conversation_id"`
	SenderId       int    `json:"sender_id"`
	ReceiverId     int    `json:"receiver_id"`
	Message        string `json:"message"`
	CreatedAt      string `json:"created_at"`
}

func NewMessage(id, conversationId, senderId, receiverId int, message, createdAt string) *Message {
	return &Message{
		Id:             id,
		ConversationId: conversationId,
		SenderId:       senderId,
		ReceiverId:     receiverId,
		Message:        message,
		CreatedAt:      createdAt,
	}
}
