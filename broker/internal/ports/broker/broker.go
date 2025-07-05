package ports

import "github.com/borismartinovic01/go-kafka/producer/internal/application/domain"

type BrokerPort interface {
	Receive(topic string, messages []*domain.Message) error
}
