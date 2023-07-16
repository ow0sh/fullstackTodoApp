package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ow0sh/fullstackTodoApp/client"
	"github.com/ow0sh/fullstackTodoApp/models"
	"github.com/ow0sh/fullstackTodoApp/usecases"
	"github.com/sirupsen/logrus"
)

type handler struct {
	httpCli *client.Client
}

func NewHandler(cli *client.Client) handler {
	return handler{httpCli: cli}
}

func (h *handler) GetLastId(ctx context.Context, todouse *usecases.TodoUseCase, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todo, err := todouse.GetLastTodo(ctx)
		if err != nil {
			log.Error(err)
		}

		jsonData, err := json.Marshal(todo.Id + 1)
		if err != nil {
			log.Error("Failed to marshal id, err: " + fmt.Sprint(err))
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func (h *handler) GetTodos(ctx context.Context, todouse *usecases.TodoUseCase, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := todouse.SelectTodos(ctx)
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

func (h *handler) InsertTodo(ctx context.Context, todouse *usecases.TodoUseCase, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var todo models.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info(fmt.Sprintf("Data recieved:text=%v, status=%v", todo.Text, todo.Status))

		todouse.CreateTodo(ctx, todo)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("todo inserted successfully"))
	}
}

func (h *handler) DeleteTodo(ctx context.Context, todouse *usecases.TodoUseCase, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		log.Info(r.Body)
		err := json.NewDecoder(r.Body).Decode(&id)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info(fmt.Sprintf("ID to delete recieved: id=%v", id))

		_, err = todouse.DeleteTodo(ctx, int64(id))
		if err != nil {
			log.Error("failed to delete todo, err: " + fmt.Sprint(err))
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("todo deleted successfully"))
	}
}

func (h *handler) SwitchStatus(ctx context.Context, todouse *usecases.TodoUseCase, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		err := json.NewDecoder(r.Body).Decode(&id)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info(fmt.Sprintf("ID to switch recieved: id=%v", id))

		todo, err := todouse.GetTodoViaId(ctx, int64(id))
		if err != nil {
			log.Error(err)
		}

		_, err = todouse.SwitchStatus(ctx, !todo.Status, int64(id))
		if err != nil {
			log.Error(err)
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("todo switched successfully"))
	}
}

func (h *handler) UpdateText(ctx context.Context, todouse *usecases.TodoUseCase, log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var request models.UpdateTextRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info(fmt.Sprintf("ID and text to updatetext recieved: id=%v", request.ID))

		_, err = todouse.UpdateTodo(ctx, request.Text, int64(request.ID))
		if err != nil {
			log.Error(err)
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("todo text updated successfully"))
	}
}
