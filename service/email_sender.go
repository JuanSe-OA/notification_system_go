// service/email_sender.go
package service

import (
    "fmt"
    "gopkg.in/gomail.v2"
    "notification-service/entity"
)

type EmailSender struct {
    Host     string
    Port     int
    Username string
    Password string
    From     string
}

func (s *EmailSender) Send(n *entity.Notification) error {
    m := gomail.NewMessage()
    m.SetHeader("From", s.From)
    m.SetHeader("To", n.Recipient)
    m.SetHeader("Subject", n.Title)
    m.SetBody("text/plain", n.Message)

    d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)

    if err := d.DialAndSend(m); err != nil {
        fmt.Printf("❌ Error enviando email: %v\n", err)
        return err
    }

    fmt.Printf("✅ Email enviado a: %s\n", n.Recipient)
    return nil
}
