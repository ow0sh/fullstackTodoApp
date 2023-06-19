package postgres

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v5"
	"github.com/ow0sh/fullstackgo/config"
)

type PSQLConn struct {
	conn *pgx.Conn
}

func NewConn(conf *config.TypePSQLConfig, log *logrus.Logger) (*PSQLConn, error) {
	config, err := pgx.ParseConfig(fmt.Sprintf("%v://%v:%v@%v:%v/%v", conf.Dsn, conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname))
	if err != nil {
		log.Error("failed to parse config")
		return nil, errors.Wrap(err, "failed to parse config")
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Error("failed to connect via config")
		return nil, errors.Wrap(err, "failed to connect via config")
	}

	return &PSQLConn{conn: conn}, nil
}

func (conn *PSQLConn) CloseConn() error {
	err := conn.conn.Close(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to close the connection")
	}
	return nil
}

type Todo struct {
	ID     int
	Text   string
	Status bool
}

func (conn *PSQLConn) InsertTodo(todo Todo) error {
	insertStr := "INSERT INTO todos VALUES ($1, $2, $3);"
	_, err := conn.conn.Exec(context.Background(), insertStr, todo.ID, todo.Text, todo.Status)
	if err != nil {
		return errors.Wrap(err, "failed to insert todo")
	}

	return nil
}

func (conn *PSQLConn) SelectLastId() (int, error) {
	var id int
	selectStr := `SELECT id FROM todos ORDER BY id DESC LIMIT 1;`
	err := conn.conn.QueryRow(context.Background(), selectStr).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to select last id")
	}
	return id, nil
}

func (conn *PSQLConn) SelectTodos() ([]Todo, error) {
	var todos []Todo

	rows, err := conn.conn.Query(context.Background(), "SELECT * FROM todos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var text string
		var status bool

		err := rows.Scan(&id, &text, &status)
		if err != nil {
			return nil, err
		}

		todos = append(todos, Todo{ID: id, Text: text, Status: status})
	}

	return todos, nil
}

func (conn *PSQLConn) DeleteTodo(id int) error {
	deleteStr := `DELETE FROM todos WHERE id=$1;`
	_, err := conn.conn.Exec(context.Background(), deleteStr, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete todo")
	}
	return nil
}

func (conn *PSQLConn) SwitchStatus(id int) error {
	var status bool
	selectStr := `SELECT status FROM todos WHERE id=$1;`
	err := conn.conn.QueryRow(context.Background(), selectStr, id).Scan(&status)
	if err != nil {
		return errors.Wrap(err, "failed to select the status")
	}

	switchStr := `UPDATE todos SET status=$1 WHERE id=$2;`
	_, err = conn.conn.Exec(context.Background(), switchStr, !status, id)
	if err != nil {
		return errors.Wrap(err, "failed to switch todo")
	}
	return nil
}

func (conn *PSQLConn) UpdateText(id int, text string) error {
	var textstr string
	selectStr := `SELECT text FROM todos WHERE id=$1;`
	err := conn.conn.QueryRow(context.Background(), selectStr, id).Scan(&textstr)
	if err != nil {
		return errors.Wrap(err, "failed to select the old text")
	}

	updateStr := `UPDATE todos SET text=$1 WHERE id=$2;`
	_, err = conn.conn.Exec(context.Background(), updateStr, text, id)
	if err != nil {
		return errors.Wrap(err, "failed to update text in todo")
	}
	return nil
}
