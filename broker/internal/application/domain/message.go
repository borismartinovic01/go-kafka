package domain

type Message struct {
	payload []byte
}

func NewMessage(payload []byte) Message {
	return Message{
		payload: payload,
	}
}
