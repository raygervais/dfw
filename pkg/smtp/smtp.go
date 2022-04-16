package smtp

import "net/smtp"

type SMTP struct {
}

func (s *SMTP) Send(to string, subject string, body string) error {

	return nil
}

// Send email using smtp package
func Send(to string, subject string, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "user", "password", "smtp.example.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail("smtp.example.com:25", auth, "", []string{to}, []byte(body))

	// Check for errors
	if err != nil {
		return err
	}

	return nil
}
