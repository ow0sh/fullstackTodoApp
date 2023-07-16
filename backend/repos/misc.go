package repos

import (
	"context"
	"database/sql"

	"github.com/ow0sh/fullstackTodoApp/models"
	"github.com/pkg/errors"
)

var (
	ErrSuchObjectAlreadyExist = errors.New("such object is already exist")
	ErrNothingUpdated         = errors.New("nothing updated")
)

type DB interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func TodosToCreateRepo(todos ...models.Todo) []CreateTodo {
	result := make([]CreateTodo, len(todos))

	for i, todo := range todos {
		result[i] = CreateTodo{
			Text: todo.Text,
		}
	}

	return result
}
