package receiver

import (
	models "messager/internal/messaging/models"
)

type MessageReceiver interface {
	ReceiveMessage() (models.Message, error)
}
