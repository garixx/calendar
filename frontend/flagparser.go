package frontend

import (
	"context"
	"log"
)

type flagParserFrontEnd struct{}

func (f flagParserFrontEnd) Start(ctx context.Context) error {
	log.Println("it's me, flag parser")
	return nil
}
