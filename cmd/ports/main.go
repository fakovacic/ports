package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fakovacic/ports/cmd/ports/config"
	"github.com/fakovacic/ports/internal/parser"
)

const errorChan int = 1

func main() {
	ctx := context.Background()

	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	store, err := config.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	service, err := config.NewService(c, store)
	if err != nil {
		log.Fatal(err)
	}

	input, err := os.Open(os.Getenv("DATA_FILENAME"))
	if err != nil {
		log.Fatalf("failed to load file: %v", err)
	}

	// validate if file is json

	parser := parser.New(service)

	errChan := make(chan error, errorChan)

	go func() {
		errChan <- parser.Parse(ctx, input)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGKILL, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				c.Log.Fatal().Msg(err.Error())
			}
		case s := <-signalChan:
			c.Log.Debug().Msgf(fmt.Sprintf("Captured %v. Exiting...", s))

			os.Exit(0)
		}
	}
}
