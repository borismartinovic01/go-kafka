package producer

import (
	"github.com/borismartinovic01/go-kafka/producer/internal/application/domain"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) Send(topic string, messages []*domain.Message) error {
	return nil
}
