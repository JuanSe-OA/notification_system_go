// service/sms_sender.go
package service

import (
    "context"
    "fmt"
    "notification_go/entity"
    "os"

    openapi "github.com/twilio/twilio-go/rest/api/v2010"
    "github.com/twilio/twilio-go"
)

type SmsSender struct {
    AccountSID string
    AuthToken  string
    FromNumber string
}

func (s *SmsSender) Send(n *entity.Notification) error {
    client := twilio.NewRestClientWithParams(twilio.ClientParams{
        Username: s.AccountSID,
        Password: s.AuthToken,
    })

    params := &openapi.CreateMessageParams{}
    params.SetTo(n.Recipient)
    params.SetFrom(s.FromNumber)
    params.SetBody(n.Message)

    _, err := client.Api.CreateMessage(params)
    if err != nil {
        fmt.Printf("❌ Error enviando SMS: %v\n", err)
        return err
    }

    fmt.Printf("✅ SMS enviado a: %s\n", n.Recipient)
    return nil
}
