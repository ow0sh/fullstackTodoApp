package sqlx

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ow0sh/fullstackTodoApp/repos"
)

type todoRepo struct {
	baseRepo[repos.Todo]
}

func NewTodoRepo(db *sqlx.DB) repos.TodoRepo {
	return &todoRepo{
		newBaseRepo[repos.Todo](db, "todos", "text", "status"),
	}
}

func (s todoRepo) Inserter() repos.TodoInserter {
	return s
}

func (s todoRepo) Create(ctx context.Context) ([]repos.Todo, error) {
	return s.baseRepo.Create(ctx)
}

func (s todoRepo) SetCreate(todos ...repos.CreateTodo) repos.TodoInserter {
	for _, todo := range todos {
		s.q.sqlInsert = s.q.sqlInsert.Values(todo.Text, false)
	}

	return s
}

func (s todoRepo) Selector() repos.TodoSelector {
	return s
}

func (s todoRepo) FilterById(ids ...int64) repos.TodoSelector {
	s.q.sqlSelect = s.q.sqlSelect.Where(squirrel.Eq{"id": ids})
	return s
}

func (s todoRepo) Limit(u uint64) repos.TodoSelector {
	s.baseRepo = s.baseRepo.Limit(u)
	return s
}

func (s todoRepo) OrderBy(by string) repos.TodoSelector {
	s.baseRepo = s.baseRepo.OrderBy(by)
	return s
}

func (s todoRepo) Deleter() repos.TodoDeleter {
	return s
}

func (s todoRepo) Delete(ctx context.Context) ([]repos.Todo, error) {
	return s.baseRepo.Delete(ctx)
}

func (s todoRepo) WhereIdD(ids ...int64) repos.TodoDeleter {
	s.q.sqlDelete = s.q.sqlDelete.Where(squirrel.Eq{"id": ids})
	return s
}

func (s todoRepo) Updater() repos.TodoUpdater {
	return s
}

func (s todoRepo) WhereIdU(ids ...int64) repos.TodoUpdater {
	s.baseRepo = s.baseRepo.WhereID(ids...)
	return s
}
