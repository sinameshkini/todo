package service

import (
	"context"
	"todo/pkg/db"
	"todo/pkg/io"
)

// TodoService describes the service.
type TodoService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Get(ctx context.Context) (t []io.Todo, error error)
	Add(ctx context.Context, todo io.Todo) (t io.Todo, error error)
	SetComplete(ctx context.Context, id string) (error error)
	RemoveComplete(ctx context.Context, id string) (error error)
	Delete(ctx context.Context, id string) (error error)
}

type basicTodoService struct{}

func (b *basicTodoService) Get(ctx context.Context) (t []io.Todo, error error) {
	session := db.ConnectPGDB()
	defer session.Close()
	error = session.Find(&t).Error
	return t, error
}
func (b *basicTodoService) Add(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	session := db.ConnectPGDB()
	defer session.Close()
	error = session.Create(&todo).Error
	return todo, error
}
func (b *basicTodoService) SetComplete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of SetComplete
	return error
}
func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of RemoveComplete
	return error
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (error error) {
	// TODO implement the business logic of Delete
	session := db.ConnectPGDB()
	defer session.Close()
	todo := io.Todo{}
	err := session.Where("id = ?", id).Find(&todo).Error
	if err != nil{
		return err
	}
	return session.Delete(&todo).Error
}

// NewBasicTodoService returns a naive, stateless implementation of TodoService.
func NewBasicTodoService() TodoService {
	return &basicTodoService{}
}

// New returns a TodoService with all of the expected middleware wired in.
func New(middleware []Middleware) TodoService {
	var svc TodoService = NewBasicTodoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
