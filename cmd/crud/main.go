package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/chelovekdanil/crud"
	// "github.com/chelovekdanil/crud/internal/config"
	"github.com/go-kit/log"
)

func main() {
	// _ = config.MustLoad()

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
		logger.Log("transport", "HTTP", "addr", ":8080")
		errs <- http.ListenAndServe(":8080", h)
	}()

	logger.Log("exit", <-errs)
}
