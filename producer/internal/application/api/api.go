package api

import (
	"strings"

	"github.com/borismartinovic01/go-kafka/producer/internal/application/domain"
	ports "github.com/borismartinovic01/go-kafka/producer/internal/ports/producer"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	producer ports.ProducerPort
}

func NewApplication(producer ports.ProducerPort) *Application {
	return &Application{
		producer: producer,
	}
}

func Send(topic string, messages []*domain.Message) {

}

func (a *Application) handleProducerSendError(err error) error {
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
		Field:       "producer",
		Description: strings.Join(allErrors, "\n"),
	}
	badReq := &errdetails.BadRequest{}
	badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
	producerStatus := status.New(codes.InvalidArgument, "message send failed")
	statusWithDetails, _ := producerStatus.WithDetails(badReq)
	return statusWithDetails.Err()
}
