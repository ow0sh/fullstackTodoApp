package sqlx

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type baseRepo[T any] struct {
	db *sqlx.DB
	q  querier
}

func newBaseRepo[T any](db *sqlx.DB, tableName string, columnsName ...string) baseRepo[T] {
	return baseRepo[T]{
		db: db,
		q:  newQuerier(tableName, columnsName...),
	}
}

func (s baseRepo[T]) Update(ctx context.Context, column string, value interface{}) ([]T, error) {
	var todos []T
	s.q.sqlUpdate = s.q.sqlUpdate.Set(column, value)
	return todos, s.q.QUpdate(ctx, &todos, s.db)
}

func (s baseRepo[T]) WhereID(ids ...int64) baseRepo[T] {
	s.q.sqlUpdate = s.q.sqlUpdate.Where(squirrel.Eq{"id": ids})
	return s
}

func (s baseRepo[T]) Limit(u uint64) baseRepo[T] {
	s.q.sqlSelect = s.q.sqlSelect.Limit(u)
	return s
}

func (s baseRepo[T]) OrderBy(by string) baseRepo[T] {
	s.q.sqlSelect = s.q.sqlSelect.OrderBy(by)
	return s
}

func (s baseRepo[T]) Get(ctx context.Context) (*T, error) {
	var result T
	return &result, s.q.QGet(ctx, &result, s.db)
}

func (s baseRepo[T]) Select(ctx context.Context) ([]T, error) {
	var result []T
	return result, s.q.QSelect(ctx, &result, s.db)
}

func (s baseRepo[T]) Create(ctx context.Context) ([]T, error) {
	var result []T
	return result, s.q.QCreate(ctx, &result, s.db)
}

func (s baseRepo[T]) Delete(ctx context.Context) ([]T, error) {
	var result []T
	return result, s.q.QDelete(ctx, &result, s.db)
}
