// entity/channel.go
package entity

type Channel string

const (
    Email    Channel = "EMAIL"
    SMS      Channel = "SMS"
    WhatsApp Channel = "WHATSAPP"
    Push     Channel = "PUSH"
)
