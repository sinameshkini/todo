package service

import (
	"context"
	"errors"
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
	Update(ctx context.Context, todo io.Todo) (t io.Todo, error error)
	SetStar(ctx context.Context, id string, star uint8) (error error)
	ReplyTo(ctx context.Context, parentId uint, todo io.Todo) (t io.Todo, error error)
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
	session := db.ConnectPGDB()
	defer session.Close()
	todo := io.Todo{}
	err := session.Where("id = ?", id).Find(&todo).Error
	if err != nil {
		return err
	}
	todo.Complete = true
	return session.Save(&todo).Error
}
func (b *basicTodoService) RemoveComplete(ctx context.Context, id string) (error error) {
	session := db.ConnectPGDB()
	defer session.Close()
	todo := io.Todo{}
	err := session.Where("id = ?", id).Find(&todo).Error
	if err != nil {
		return err
	}
	todo.Complete = false
	return session.Save(&todo).Error
}
func (b *basicTodoService) Delete(ctx context.Context, id string) (error error) {
	session := db.ConnectPGDB()
	defer session.Close()
	todo := io.Todo{}
	err := session.Where("id = ?", id).Find(&todo).Error
	if err != nil {
		return err
	}
	return session.Delete(&todo).Error
}

func (b *basicTodoService) Update(ctx context.Context, todo io.Todo) (t io.Todo, error error) {
	session := db.ConnectPGDB()
	defer session.Close()
	error = session.Save(&todo).Error
	return todo, error
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

func (b *basicTodoService) SetStar(ctx context.Context, id string, star uint8) (error error) {
	session := db.ConnectPGDB()
	defer session.Close()
	todo := io.Todo{}
	err := session.Where("id = ?", id).Find(&todo).Error
	if err != nil {
		return err
	}
	if star < 0 || star > 5 {
		return errors.New("star value out of range. valid range is 0 to 5")
	}
	todo.Star = star
	return session.Save(&todo).Error
}

func (b *basicTodoService) ReplyTo(ctx context.Context, parentId uint, todo io.Todo) (t io.Todo, error error) {
	// TODO implement the business logic of ReplyTo
	session := db.ConnectPGDB()
	defer session.Close()
	todo.ParentID = parentId
	error = session.Create(&todo).Error
	return todo, error
}
