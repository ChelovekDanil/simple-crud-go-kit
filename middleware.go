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

func (mw loggerMiddleware) GetAll(ctx context.Context) (users []User, err error) {
	defer func(t time.Time) {
		mw.logger.Log("method", "GetAll", "time", time.Since(t), "error", err)
	}(time.Now())
	return mw.next.GetAll(ctx)
}

func (mw loggerMiddleware) Create(ctx context.Context, user User) (id string, err error) {
	defer func(t time.Time) {
		mw.logger.Log("method", "Create", "first name", user.FirstName, "last name", user.LastName, "time", time.Since(t), "error", err)
	}(time.Now())
	return mw.next.Create(ctx, user)
}

func (mw loggerMiddleware) Update(ctx context.Context, user User) (err error) {
	defer func(t time.Time) {
		mw.logger.Log("method", "update", "id", user.Id, "first name", user.FirstName, "last name", user.LastName, "time", time.Since(t), "error", err)
	}(time.Now())
	return mw.next.Update(ctx, user)
}

func (mw loggerMiddleware) Delete(ctx context.Context, id string) (err error) {
	defer func(t time.Time) {
		mw.logger.Log("method", "delete", "id", id, "time", time.Since(t), "error", err)
	}(time.Now())
	return mw.next.Delete(ctx, id)
}
