package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/pkg/errors"

	"github.com/ow0sh/fullstackTodoApp/client"
	"github.com/ow0sh/fullstackTodoApp/config"
	"github.com/ow0sh/fullstackTodoApp/middlware"
	sqlx2 "github.com/ow0sh/fullstackTodoApp/repos/sqlx"
	"github.com/ow0sh/fullstackTodoApp/usecases"
)

const configPath = "./config.json"

func main() {
	config, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	log := config.Log()
	db := config.DB()

	ctx, cancel := ctxWithSig()
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
			cancel()
		}
	}()

	httpcli := &http.Client{}
	client := client.NewClient(httpcli)

	handler := NewHandler(client)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(middlware.Logger(log))

	todosUse := usecases.NewTodoUseCase(sqlx2.NewTodoRepo(db))

	r.Route("/api", func(r chi.Router) {
		r.Get("/getlastid", handler.GetLastId(ctx, todosUse, log))
		r.Post("/inserttodo", handler.InsertTodo(ctx, todosUse, log))
		r.Get("/gettodos", handler.GetTodos(ctx, todosUse, log))
		r.Delete("/deletetodo", handler.DeleteTodo(ctx, todosUse, log))
		r.Post("/switchstatus", handler.SwitchStatus(ctx, todosUse, log))
		r.Post("/updatetext", handler.UpdateText(ctx, todosUse, log))
	})

	log.Info("Server has been started")
	if err := http.ListenAndServe(":3001", r); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}

func ctxWithSig() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	go func() {
		select {
		case <-ch:
			cancel()
		}
	}()

	return ctx, cancel
}
