package processor

import (
	"log"
	msg "messenger/internal/messaging/models/message"

	"github.com/gorilla/websocket"
)

type WebSocketMessageProcessor struct {
	connection          *websocket.Conn
	errorResponseText   string
	infoResponseText    string
	dataResponseText    string
	unknownResponseText string
}

type Options struct {
	ErrorResponseText   string
	InfoResponseText    string
	DataResponseText    string
	UnknownResponseText string
}

// New создает новый экземпляр WebSocketMessageProcessor с предоставленными параметрами.
// Параметр options позволяет настроить тексты ответов для ошибок, информационных сообщений и данных.
//
// Параметры:
//   - options: Структура Options, содержащая конфигурацию для WebSocketMessageProcessor.
//
// Возвращает:
//
//	Указатель на вновь инициализированный WebSocketMessageProcessor.
func New(options Options) *WebSocketMessageProcessor {
	return &WebSocketMessageProcessor{
		connection:          nil,
		errorResponseText:   options.ErrorResponseText,
		infoResponseText:    options.InfoResponseText,
		dataResponseText:    options.DataResponseText,
		unknownResponseText: options.UnknownResponseText,
	}
}

// SetConnection устанавливает WebSocket-соединение для WebSocketMessageProcessor.
//
// Параметры:
//   - conn: Указатель на websocket.Conn, представляющий устанавливаемое WebSocket-соединение.
func (wsmp *WebSocketMessageProcessor) SetConnection(conn *websocket.Conn) {
	wsmp.connection = conn
}

// ProcessMessage обрабатывает входящее сообщение WebSocket в зависимости от его типа.
// Обрабатывает различные типы сообщений ("error", "info", "data") с использованием соответствующих
// методов обработки и возвращает ответное сообщение.
//
// Параметры:
//   - message: Входящее сообщение типа msg.Message для обработки.
//
// Возвращает:
//   - msg.Message: Обработанное ответное сообщение.
//   - error: Ошибка, если WebSocket-соединение не установлено или возникла другая проблема.
//
// Поведение:
//   - Если WebSocket-соединение не установлено, возвращает ошибку типа *websocket.CloseError.
//   - Для известных типов сообщений (ErrorMessage, InfoMessage, DataMessage) обрабатывает сообщение с использованием
//     соответствующих методов (processError, processInfo, processData).
//   - Для неизвестных типов сообщений регистрирует проблему и возвращает ответное сообщение
//     с типом UnknownResponse и описанием ошибки.
func (wsmp *WebSocketMessageProcessor) ProcessMessage(message msg.Message) (msg.Message, error) {
	if wsmp.connection != nil {
		switch message.Type {
		case msg.ErrorMessage:
			responseMessage := wsmp.processError(message, wsmp.errorResponseText)
			return responseMessage, nil
		case msg.InfoMessage:
			responseMessage := wsmp.processInfo(message, wsmp.infoResponseText)
			return responseMessage, nil
		case msg.DataMessage:
			responseMessage := wsmp.processData(message, wsmp.dataResponseText)
			return responseMessage, nil
		default:
			log.Printf("Получен неизвестный тип сообщения: %s\n", message)
			responseMessage := msg.Message{
				Type: msg.UnknownResponse,
				Text: "Неизвестный тип сообщения",
			}
			return responseMessage, nil
		}
	} else {
		return msg.Message{}, &websocket.CloseError{Code: websocket.CloseAbnormalClosure, Text: "Соединение не установлено"}
	}
}

// createResponseMessage создает новое ответное сообщение с указанным типом сообщения и текстом.
//
// Параметры:
//   - messageType: Тип создаваемого сообщения.
//   - responseText: Текстовое содержимое ответного сообщения.
//
// Возвращает:
//   - Экземпляр msg.Message, содержащий указанный тип и текст.
func (wsmp *WebSocketMessageProcessor) createResponseMessage(messageType msg.MessageType, responseText string) msg.Message {
	return msg.Message{
		Type: messageType,
		Text: responseText,
	}
}

// processError обрабатывает сообщение об ошибке, полученное от клиента.
// Регистрирует сообщение об ошибке и возвращает ответное сообщение, указывающее,
// что сообщение об ошибке было получено.
//
// Параметры:
//   - errorMessage: Сообщение об ошибке, отправленное клиентом.
//   - responseText: Предопределенный текст для ответа.
//
// Возвращает:
//   - msg.Message: Ответное сообщение с типом "error_response" и предоставленным текстом ответа.
func (wsmp *WebSocketMessageProcessor) processError(errorMessage msg.Message, responseText string) msg.Message {
	log.Printf("Клиент отправил сообщение об ошибке: %s\n", errorMessage.Text)
	return wsmp.createResponseMessage(msg.ErrorMessage, responseText)
}

// processInfo обрабатывает информационное сообщение, полученное от клиента.
// Регистрирует содержимое сообщения и возвращает ответное сообщение,
// подтверждающее получение информационного сообщения.
//
// Параметры:
//   - infoMessage: Информационное сообщение, отправленное клиентом.
//   - responseText: Предопределенный текст для ответа.
//
// Возвращает:
//   - msg.Message: Ответное сообщение, указывающее, что информационное
//     сообщение было получено.
func (wsmp *WebSocketMessageProcessor) processInfo(infoMessage msg.Message, responseText string) msg.Message {
	log.Printf("Клиент отправил информационное сообщение: %s\n", infoMessage.Text)
	return wsmp.createResponseMessage(msg.InfoResponse, responseText)
}

// processData обрабатывает входящее сообщение с данными и генерирует ответное сообщение.
// Регистрирует текст полученного сообщения и возвращает предопределенный ответ.
//
// Параметры:
//   - dataMessage: Входящее сообщение типа msg.Message, содержащее данные.
//   - responseText: Предопределенный текст для ответа.
//
// Возвращает:
//   - msg.Message: Ответное сообщение с типом "data_response" и предоставленным текстом ответа.
func (wsmp *WebSocketMessageProcessor) processData(dataMessage msg.Message, responseText string) msg.Message {
	log.Printf("Клиент отправил сообщение с данными: %s\n", dataMessage.Text)
	return wsmp.createResponseMessage(msg.DataResponse, responseText)
}
