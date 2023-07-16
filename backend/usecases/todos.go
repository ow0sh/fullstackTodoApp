package usecases

import (
	"context"
	"database/sql"

	"github.com/ow0sh/fullstackTodoApp/models"
	"github.com/ow0sh/fullstackTodoApp/repos"
	"github.com/pkg/errors"
)

type TodoUseCase struct {
	repo repos.TodoRepo
}

func NewTodoUseCase(repo repos.TodoRepo) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (use TodoUseCase) SelectTodos(ctx context.Context) ([]repos.Todo, error) {
	todos, err := use.repo.Selector().Select(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select todos")
	}

	return todos, nil
}

func (use TodoUseCase) CreateTodo(ctx context.Context, todos ...models.Todo) ([]repos.Todo, error) {
	todosDB := make([]repos.Todo, 0, len(todos))
	for _, todo := range todos {
		todosDB, err := use.repo.Inserter().SetCreate(repos.TodosToCreateRepo(todo)...).Create(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create todo")
		}
		todosDB = append(todosDB, todosDB[0])
	}

	return todosDB, nil
}

func (use TodoUseCase) UpdateTodo(ctx context.Context, text string, id int64) ([]repos.Todo, error) {
	todosDB, err := use.repo.Updater().WhereIdU(id).Update(ctx, "text", text)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update todo")
	}

	return todosDB, nil
}

func (use TodoUseCase) DeleteTodo(ctx context.Context, id int64) ([]repos.Todo, error) {
	todosDB, err := use.repo.Deleter().WhereIdD(id).Delete(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete todo")
	}

	return todosDB, nil
}

func (use TodoUseCase) SwitchStatus(ctx context.Context, status bool, id int64) ([]repos.Todo, error) {
	todosDB, err := use.repo.Updater().WhereIdU(id).Update(ctx, "status", status)
	if err != nil {
		return nil, errors.Wrap(err, "failed to switch status")
	}

	return todosDB, err
}

func (use TodoUseCase) GetTodoViaId(ctx context.Context, id int64) (*repos.Todo, error) {
	todo, err := use.repo.Selector().FilterById(id).Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get todo via id")
	}

	return todo, nil
}

func (use TodoUseCase) GetLastTodo(ctx context.Context) (*repos.Todo, error) {
	todo, err := use.repo.Selector().Limit(1).OrderBy("id").Get(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &repos.Todo{Id: 1, Text: "", Status: false}, nil
		}
		return nil, errors.Wrap(err, "failed to get last todo")
	}

	return todo, nil
}
