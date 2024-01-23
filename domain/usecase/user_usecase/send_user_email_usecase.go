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

func (cs *SendAuthEmailUseCase) Execute(userId int, senderEmail string, receiverEmail string, subject string, htmlBody string, passcode string) error {
	if cs == nil {
		fmt.Printf("Email service error: SendAuthEmailUseCase is nil")
		return fmt.Errorf("Email service error")
	}

	_, errr := cs.userRepo.UpdateUserEmailPasscode(userId, passcode)

	if errr != nil {
		fmt.Printf("updating passcode on db err: %v", errr)
		return fmt.Errorf("update user email passcode error")
	}

	err := cs.emailRepo.SendEmail(senderEmail, receiverEmail, subject, htmlBody)
	if err != nil {
		fmt.Printf("Email service error: %v", err)
		return fmt.Errorf("Email service error")
	}

	return nil
}
