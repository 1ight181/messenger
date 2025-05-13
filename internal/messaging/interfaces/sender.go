package interfaces

import (
	message "messenger/internal/messaging/models/message"
)

type MessageSender interface {
	SendMessage(message message.Message) error
}
