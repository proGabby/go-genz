package email_repo

type EmailSender interface {
	SendEmail(senderEmail string, receiverEmail string, subject string, htmlBody string) error
}
