package broker

import "github.com/borismartinovic01/go-kafka/producer/internal/application/domain"

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Receive(topic string, messages []*domain.Message) error {
	return nil
}
