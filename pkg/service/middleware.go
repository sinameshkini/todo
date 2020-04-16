package service

import (
	"context"
	io "todo/pkg/io"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(TodoService) TodoService

type loggingMiddleware struct {
	logger log.Logger
	next   TodoService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a TodoService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next TodoService) TodoService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Get(ctx context.Context) (t []io.Todo, error error) {
	defer func() {
		l.logger.Log("method", "Get", "t", t, "error", error)
	}()
	return l.next.Get(ctx)
}
func (l loggingMiddleware) Add(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	defer func() {
		l.logger.Log("method", "Add", "todo", todo, "t", t, "error", error)
	}()
	return l.next.Add(ctx, todo)
}
func (l loggingMiddleware) SetComplete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "SetComplete", "id", id, "error", error)
	}()
	return l.next.SetComplete(ctx, id)
}
func (l loggingMiddleware) RemoveComplete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "RemoveComplete", "id", id, "error", error)
	}()
	return l.next.RemoveComplete(ctx, id)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (error error) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "error", error)
	}()
	return l.next.Delete(ctx, id)
}

func (l loggingMiddleware) Update(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	defer func() {
		l.logger.Log("method", "Update", "todo", todo, "t", t, "error", error)
	}()
	return l.next.Update(ctx, todo)
}

func (l loggingMiddleware) SetStar(ctx context.Context, id string, star uint8) (error error) {
	defer func() {
		l.logger.Log("method", "SetStar", "id", id, "star", star, "error", error)
	}()
	return l.next.SetStar(ctx, id, star)
}

func (l loggingMiddleware) ReplyTo(ctx context.Context, parentId uint, todo io.Todo) (t io.Todo, error error) {
	defer func() {
		l.logger.Log("method", "ReplyTo", "parentId", parentId, "todo", todo, "t", t, "error", error)
	}()
	return l.next.ReplyTo(ctx, parentId, todo)
}
