package frontend

import (
	"context"
	"log"
)

type cliFrontEnd struct{}

func (f cliFrontEnd) Start(ctx context.Context) error {
	log.Println("it's me, cli frontend")
	return nil
}
