package chatusecase

import (
	"github.com/proGabby/4genz/domain/entity"
	"github.com/proGabby/4genz/domain/repository/chat_repo"
)

type SaveChatMsgUseCase struct {
	chatRepo chat_repo.ChatRepository
}

func NewSaveChatMsgUseCase(chatRepo chat_repo.ChatRepository) *SaveChatMsgUseCase {
	return &SaveChatMsgUseCase{
		chatRepo: chatRepo,
	}
}

func (saveChatMsgUseCase *SaveChatMsgUseCase) Execute(message *entity.Message) error {
	return saveChatMsgUseCase.chatRepo.SaveMessage(message)
}
