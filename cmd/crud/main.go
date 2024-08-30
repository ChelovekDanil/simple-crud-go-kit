package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/chelovekdanil/crud"
	"github.com/go-kit/log"
)

const (
	defaultServerPort = "8080"
	PORT_ENV          = "PORT"
)

func main() {
	var port string
	if port = os.Getenv(PORT_ENV); port == "" {
		port = defaultServerPort
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var service crud.Service
	{
		service = crud.NewService()
		service = crud.LoggerMiddleware(logger)(service)
	}

	var h http.Handler
	{
		h = crud.MakeHTTPHandler(service, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", ":"+port)
		errs <- http.ListenAndServe(":"+port, h)
	}()

	logger.Log("exit", <-errs)
}
