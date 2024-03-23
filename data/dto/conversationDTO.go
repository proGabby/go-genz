package dto

type ConversationDTO struct {
	Id        int    `json:"id"`
	RoomId string `json:"room_id"`
	InitiatorId  int    `json:"initiator_id"`
	RecipientId int    `json:"recipient_id"`
	ConnectedAt string `json:"connected_at"`
}