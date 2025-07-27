package notifier

import "fmt"

type Notifier interface {
	Send(to, message string) error
}

type EmailService struct{}

func (e *EmailService) Send(to, message string) error {
	fmt.Printf("Sending email to %s: %s\n", to, message)
	return nil
}

func NotifyUser(n Notifier, user string) error {
	if user == "" {
		return fmt.Errorf("no user specified")
	}
	return n.Send(user, "Welcome to Djamware!")
}
