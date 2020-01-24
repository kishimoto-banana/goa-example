package exampleapi

import (
	"context"
	ping "goa-example/gen/ping"
	"log"
)

// ping service example implementation.
// The example methods log the requests and return zero values.
type pingsrvc struct {
	logger *log.Logger
}

// NewPing returns the ping service implementation.
func NewPing(logger *log.Logger) ping.Service {
	return &pingsrvc{logger}
}

// ReturnPong implements Return pong.
func (s *pingsrvc) ReturnPong(ctx context.Context) (res string, err error) {
	s.logger.Print("ping.Return pong")
	return "pong", nil
}
