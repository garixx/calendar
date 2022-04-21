package frontend

import (
	"context"
	"fmt"
	"log"
)

type Frontend interface {
	Start(ctx context.Context) error
}

func NewFrontEnd(s string) (Frontend, error) {
	switch s {
	case "null", "nil":
		return nullFrontEnd{}, nil
	case "cli":
		return &cliFrontEnd{}, nil
	case "flagparser":
		return &flagParserFrontEnd{}, nil
	case "rest":
		return &restFrontEnd{}, nil
	default:
		return nil, fmt.Errorf("no such frontend %s", s)
	}
}

type nullFrontEnd struct{}

func (f nullFrontEnd) Start(ctx context.Context) error {
	log.Println("it's me, null frontend")
	return nil
}
