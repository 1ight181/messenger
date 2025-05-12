package interfaces

import (
	models "messenger/internal/messaging/models"
)

type MessageReceiver interface {
	ReceiveMessage() (models.Message, error)
}
