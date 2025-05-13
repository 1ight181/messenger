package interfaces

import (
	message "messenger/internal/messaging/models/message"
)

type MessageReceiver interface {
	ReceiveMessage() (message.Message, error)
}
