package main

import "fmt"

type Mailer interface {
	Send(to, body string) error
}

type SMTPMailer struct{}

func (m SMTPMailer) Send(to, body string) error {
	fmt.Printf("mail sent to: %s\n", to)
	return nil
}

type MockMailer struct{}

func (m MockMailer) Send(to, body string) error {
	fmt.Printf("[Mock] mail sent to: %s\n", to)
	return nil
}

func Notify(m Mailer, to string) {
	m.Send(to, "hello")
}

func main() {
	Notify(SMTPMailer{}, "user@example.com")
	Notify(MockMailer{}, "user@example.com")
}
