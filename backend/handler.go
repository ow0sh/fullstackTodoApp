package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ow0sh/fullstackgo/client"
	"github.com/ow0sh/fullstackgo/postgres"
	"github.com/sirupsen/logrus"
)

type handler struct {
	httpCli *client.Client
}

func NewHandler(cli *client.Client) handler {
	return handler{httpCli: cli}
}

func (h *handler) getId(conn *postgres.PSQLConn, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := conn.SelectLastId()
		if err != nil {
			log.Error(err)
		}

		jsonData, err := json.Marshal(id)
		if err != nil {
			log.Error("Failed to marshal id, err: " + fmt.Sprint(err))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (h *handler) GetTodos(conn *postgres.PSQLConn, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := conn.SelectTodos()
		if err != nil {
			log.Error("Failed to select todos, err: " + fmt.Sprint(err))

		}

		jsonData, err := json.Marshal(todos)
		if err != nil {
			log.Error("Failed to marshal todos, err: " + fmt.Sprint(err))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (h *handler) InsertTodo(conn *postgres.PSQLConn, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo postgres.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info(fmt.Sprintf("Data recieved: id=%v, text=%v, status=%v", todo.ID, todo.Text, todo.Status))

		conn.InsertTodo(todo)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("todo inserted successfully"))
	}
}

func (h *handler) DeleteTodo(conn *postgres.PSQLConn, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		err := json.NewDecoder(r.Body).Decode(&id)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info(fmt.Sprintf("ID to delete recieved: id=%v", id))

		conn.DeleteTodo(id)

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("todo deleted successfully"))
	}
}
