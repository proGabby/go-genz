package chatusecase

type ChatUsecases struct {
	SaveChatMsg SaveChatMsgUseCase
}

func NewChatUsecases(saveChatMsg SaveChatMsgUseCase) *ChatUsecases {
	return &ChatUsecases{
		SaveChatMsg: saveChatMsg,
	}
}
