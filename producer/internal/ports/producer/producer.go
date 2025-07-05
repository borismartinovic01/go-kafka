package ports

import "github.com/borismartinovic01/go-kafka/producer/internal/application/domain"

type ProducerPort interface {
	Send(topic string, messages []*domain.Message) error
}
