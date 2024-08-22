package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-app/pkg/entities"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
}

type TodoList interface {
	Create(userId int, list entities.TodoList) (int, error)
	GetAll(userId int) ([]entities.TodoList, error)
	GetById(userId, listId int) (entities.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input entities.UpdateListInput) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db)}
}
