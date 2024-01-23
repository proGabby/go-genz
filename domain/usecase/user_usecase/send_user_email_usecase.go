package user_usecase

import (
	"fmt"

	email_repo "github.com/proGabby/4genz/domain/repository/email_handler_repo"
	"github.com/proGabby/4genz/domain/repository/user_repo"
)

type SendAuthEmailUseCase struct {
	userRepo  user_repo.UserRepository
	emailRepo email_repo.EmailSender
}

func NewSendAuthEmailUseCase(userRepo user_repo.UserRepository, emailRepo email_repo.EmailSender) *SendAuthEmailUseCase {

	return &SendAuthEmailUseCase{
		userRepo:  userRepo,
		emailRepo: emailRepo,
	}
}

func (cs *SendAuthEmailUseCase) Execute(senderEmail string, receiverEmail string, subject string, htmlBody string) error {
	if cs == nil {
		fmt.Printf("Email service error: SendAuthEmailUseCase is nil")
		return fmt.Errorf("Email service error")
	}
	return cs.emailRepo.SendEmail(senderEmail, receiverEmail, subject, htmlBody)
}
