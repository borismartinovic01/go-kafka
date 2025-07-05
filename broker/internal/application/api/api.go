package api

import (
	"strings"

	"github.com/borismartinovic01/go-kafka/producer/internal/application/domain"
	ports "github.com/borismartinovic01/go-kafka/producer/internal/ports/broker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	broker ports.BrokerPort
}

func NewApplication(broker ports.BrokerPort) *Application {
	return &Application{
		broker: broker,
	}
}

func (a *Application) Receive(topic string, messages []*domain.Message) error {
	err := a.broker.Receive(topic, messages)
	if err != nil {
		return a.handleBrokerReceiveError(err)
	}
	return nil
}

func (a *Application) handleBrokerReceiveError(err error) error {
	st := status.Convert(err)
	var allErrors []string
	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.BadRequest:
			for _, violation := range t.GetFieldViolations() {
				allErrors = append(allErrors, violation.Description)
			}
		}
	}
	fieldErr := &errdetails.BadRequest_FieldViolation{
		Field:       "broker",
		Description: strings.Join(allErrors, "\n"),
	}
	badReq := &errdetails.BadRequest{}
	badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
	producerStatus := status.New(codes.InvalidArgument, "message send failed")
	statusWithDetails, _ := producerStatus.WithDetails(badReq)
	return statusWithDetails.Err()
}
