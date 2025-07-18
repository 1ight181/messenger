package interfaces

import (
	message "messenger/internal/messaging/models/message"
)

type MessageProcessor interface {
	ProcessMessage(message.Message) (message.Message, error)
}
