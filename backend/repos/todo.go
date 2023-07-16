package repos

import "context"

type TodoSelector interface {
	FilterById(...int64) TodoSelector
	OrderBy(string) TodoSelector
	Limit(uint64) TodoSelector

	Select(ctx context.Context) ([]Todo, error)
	Get(ctx context.Context) (*Todo, error)
}

type TodoInserter interface {
	SetCreate(...CreateTodo) TodoInserter
	Create(ctx context.Context) ([]Todo, error)
}

type TodoUpdater interface {
	WhereIdU(...int64) TodoUpdater
	Update(ctx context.Context, column string, value interface{}) ([]Todo, error)
}

type TodoDeleter interface {
	WhereIdD(...int64) TodoDeleter
	Delete(ctx context.Context) ([]Todo, error)
}

type TodoRepo interface {
	Updater() TodoUpdater
	Inserter() TodoInserter
	Selector() TodoSelector
	Deleter() TodoDeleter
}

type CreateTodo struct {
	Text string `db:"text"`
}

type Todo struct {
	Id     int64  `db:"id"`
	Text   string `db:"text"`
	Status bool   `db:"status"`
}
