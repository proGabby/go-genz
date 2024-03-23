package entity

type Message struct {
	Id         int    `json:"id"`
	RoomId     string `json:"room_id"`
	SenderId   int    `json:"sender_id"`
	ReceiverId int    `json:"receiver_id"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
}
