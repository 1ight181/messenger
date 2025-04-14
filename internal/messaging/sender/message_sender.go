package sender

import (
	messaging "messager/internal/messaging/types"
)

type MessageSender interface {
	SendMessage(message messaging.Message) error
}
