package receiver

import (
	messaging "messager/internal/messaging/types"
)

type MessageReceiver interface {
	ReceiveMessage() (messaging.Message, error)
}
