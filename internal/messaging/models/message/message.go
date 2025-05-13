package message

type Message struct {
	Type MessageType `json:"type"`
	Text string      `json:"text"`
}
