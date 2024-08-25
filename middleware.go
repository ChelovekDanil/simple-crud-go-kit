package crud

import (
	"context"
	"time"

	"github.com/go-kit/log"
)

type Middleware func(Service) Service

func LoggerMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggerMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggerMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggerMiddleware) Get(ctx context.Context, id string) (user *User, err error) {
	defer func(t time.Time) {
		mw.logger.Log("method", "Get", "id", id, "time", time.Since(t), "error", err)
	}(time.Now())
	return mw.next.Get(ctx, id)
}
