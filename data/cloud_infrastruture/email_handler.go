package infrastruture

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

type GomailSender struct {
	emailConfig   *gomail.Message
	dialerCreator *gomail.Dialer
}

func NewGomailSender(host string) (*GomailSender, error) {
	gPassword, ok := os.LookupEnv("GOOGLE_APP_PASSWORD");
	if !ok {
		fmt.Printf("Google app password not found on env")
		return nil, fmt.Errorf("Google app password not found on env")
	}
	gEmail, ok := os.LookupEnv("GOOGLE_APP_EMAIL");
	if !ok {
		fmt.Printf("Google app email not found on env")
		return nil, fmt.Errorf("Google app email not found on env")
	}

	return &GomailSender{
		emailConfig:   gomail.NewMessage(),
		dialerCreator: gomail.NewDialer(host, 465, gEmail, gPassword),
	}, nil
}

func (gmtp *GomailSender) SendEmail(senderEmail string, receiverEmail string, subject string, htmlBody string) error {

	gmtp.emailConfig.SetHeader("From", senderEmail)
	gmtp.emailConfig.SetHeader("To", receiverEmail)
	gmtp.emailConfig.SetHeader("Subject", subject)

	gmtp.emailConfig.SetBody("text/html", htmlBody)

	// Create a new SMTP dialer
	// dialer := gomail.NewDialer("smtp.gmail.com", 587, "your-email@gmail.com", "your-email-password")

	// Send the email using the dialer
	err := gmtp.dialerCreator.DialAndSend(gmtp.emailConfig)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
