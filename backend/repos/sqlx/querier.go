package sqlx

import (
	"context"
	sqlerr "database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/ow0sh/fullstackTodoApp/repos"
	"github.com/pkg/errors"
)

type querier struct {
	sqlSelect squirrel.SelectBuilder
	sqlInsert squirrel.InsertBuilder
	sqlUpdate squirrel.UpdateBuilder
	sqlDelete squirrel.DeleteBuilder

	table string
}

func newQuerier(table string, insertColumns ...string) querier {
	return querier{
		table:     table,
		sqlSelect: squirrel.Select(fmt.Sprintf("%s.*", table)).From(fmt.Sprintf("%s", table)).PlaceholderFormat(squirrel.Dollar),
		sqlInsert: squirrel.Insert(table).Columns(insertColumns...).PlaceholderFormat(squirrel.Dollar),
		sqlUpdate: squirrel.Update(table).PlaceholderFormat(squirrel.Dollar),
		sqlDelete: squirrel.Delete(table).PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *querier) QUpdate(ctx context.Context, dest interface{}, q repos.DB) error {
	stmt := r.sqlUpdate.Suffix("RETURNING *")

	sql, args, err := stmt.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to convert to sql")
	}

	err = q.SelectContext(ctx, dest, sql, args...)
	if err != nil {
		return errors.Wrap(err, "failed to exec stmt")
	}

	return nil
}

func (r *querier) QGet(ctx context.Context, dest interface{}, q repos.DB) error {
	sql, args, err := r.sqlSelect.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to convert to sql")
	}

	err = q.GetContext(ctx, dest, sql, args...)
	if err != nil {
		if errors.Is(err, sqlerr.ErrNoRows) {
			return sqlerr.ErrNoRows
		}
		return errors.Wrap(err, "failed to exec get stmt")
	}

	return nil
}

func (r *querier) QSelect(ctx context.Context, dest interface{}, q repos.DB) error {
	sql, args, err := r.sqlSelect.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to convert to sql")
	}

	err = q.SelectContext(ctx, dest, sql, args...)
	if err != nil {
		if errors.Is(err, sqlerr.ErrNoRows) {
			return sqlerr.ErrNoRows
		}
		return errors.Wrap(err, "failed to exec select stmt")
	}

	return nil
}

func (r *querier) QCreate(ctx context.Context, dest interface{}, q repos.DB) error {
	stmt := r.sqlInsert.Suffix("RETURNING *")

	sql, args, err := stmt.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to convert to sql")
	}

	err = q.SelectContext(ctx, dest, sql, args...)
	if err != nil {
		return errors.Wrap(err, "failed to exec stmt")
	}

	return nil
}

func (r *querier) QDelete(ctx context.Context, dest interface{}, q repos.DB) error {
	sql, args, err := r.sqlDelete.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to convert to sql")
	}

	_, err = q.ExecContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}

	return nil
}
