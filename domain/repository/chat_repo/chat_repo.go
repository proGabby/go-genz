package chat_repo

import "github.com/proGabby/4genz/domain/entity"

type ChatRepository interface {
	SaveMessage(message *entity.Message) error
	InitConversation(senderId, receiverId int) (int, error)
	GetUserConversations(userId int) ([]entity.Conversation, error)
	GetMessagesByConversationId(conversationId int) ([]entity.Message, error)
}
