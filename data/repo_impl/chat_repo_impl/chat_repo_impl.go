package chat_repo_impl

import (
	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	"github.com/proGabby/4genz/domain/entity"
)

type ChatRepositoryImpl struct {
	psql postgressDatasource.PostgresChatDBStore
}

func NewChatRepoImpl(psql postgressDatasource.PostgresChatDBStore) *ChatRepositoryImpl {
	return &ChatRepositoryImpl{
		psql: psql,
	}
}

func (chatRepoImpl *ChatRepositoryImpl) SaveMessage(message *entity.Message) error {
	return chatRepoImpl.psql.SaveMessage(message)
}

func (chatRepoImpl *ChatRepositoryImpl) InitConversation(senderId, receiverId int) (int, error) {
	// return chatRepoImpl.psql.InitConversation(senderId, receiverId)
	panic("not implemented")
}

func (chatRepoImpl *ChatRepositoryImpl) GetUserConversations(userId int) ([]entity.Conversation, error) {
	// return chatRepoImpl.psql.GetUserConversations(userId)
	panic("not implemented")
}

func (chatRepoImpl *ChatRepositoryImpl) GetMessagesByConversationId(conversationId int) ([]entity.Message, error) {
	// return chatRepoImpl.psql.GetMessagesByConversationId(conversationId)
	panic("not implemented")
}
