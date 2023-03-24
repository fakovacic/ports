package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/fakovacic/ports/cmd/ports/config"
	"github.com/fakovacic/ports/internal/ports/handlers/http/middleware"
	"github.com/julienschmidt/httprouter"
)

const errorChan int = 1

func main() {
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

	h := config.NewHandlers(c, service)

	router := httprouter.New()

	router.POST("/ports", h.Create)
	router.PUT("/ports/:id", h.Update)

	var (
		httpAddr = "0.0.0.0:8080"
	)

	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
	httpServer.Handler = middleware.ReqID(
		middleware.Logger(
			c, router,
		),
	)

	errChan := make(chan error, errorChan)

	go func() {
		c.Log.Debug().Msgf("HTTP service listening on %s", httpAddr)
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				c.Log.Fatal().Msg(err.Error())
			}
		case s := <-signalChan:
			c.Log.Debug().Msgf(fmt.Sprintf("Captured %v. Exiting...", s))

			httpServer.BlockingClose()
			os.Exit(0)
		}
	}

}
