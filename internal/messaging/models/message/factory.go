package message

// NewErrorMessage создает сообщение с типом ErrorMessage и заданным текстом.
func NewErrorMessage(text string) Message {
	return Message{
		Type: ErrorMessage,
		Text: text,
	}
}

// NewInfoMessage создает сообщение с типом InfoMessage и заданным текстом.
func NewInfoMessage(text string) Message {
	return Message{
		Type: InfoMessage,
		Text: text,
	}
}

// NewDataMessage создает сообщение с типом DataMessage и заданным текстом.
func NewDataMessage(text string) Message {
	return Message{
		Type: DataMessage,
		Text: text,
	}
}

// NewErrorResponse создает ответное сообщение с типом ErrorResponse и заданным текстом.
func NewErrorResponse(text string) Message {
	return Message{
		Type: ErrorResponse,
		Text: text,
	}
}

// NewInfoResponse создает ответное сообщение с типом InfoResponse и заданным текстом.
func NewInfoResponse(text string) Message {
	return Message{
		Type: InfoResponse,
		Text: text,
	}
}

// NewDataResponse создает ответное сообщение с типом DataResponse и заданным текстом.
func NewDataResponse(text string) Message {
	return Message{
		Type: DataResponse,
		Text: text,
	}
}

// NewUnknownResponse создает ответное сообщение с типом UnknownResponse и заданным текстом.
func NewUnknownResponse(text string) Message {
	return Message{
		Type: UnknownResponse,
		Text: text,
	}
}
