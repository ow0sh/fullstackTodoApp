package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/ow0sh/fullstackgo/client"
	"github.com/ow0sh/fullstackgo/config"
	"github.com/ow0sh/fullstackgo/middlware"
	"github.com/ow0sh/fullstackgo/postgres"
)

func main() {
	log := logrus.New()

	config, _ := config.InitConfig(log)

	httpcli := &http.Client{}
	client := client.NewClient(httpcli)

	handler := NewHandler(client)

	PSQLConn, _ := postgres.NewConn(&config.PSQL, log)

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

	r.Route("/api", func(r chi.Router) {
		r.Get("/getlastid", handler.getId(PSQLConn, log))
		r.Post("/inserttodo", handler.InsertTodo(PSQLConn, log))
		r.Get("/gettodos", handler.GetTodos(PSQLConn, log))
		r.Delete("/deletetodo", handler.DeleteTodo(PSQLConn, log))
	})

	if err := http.ListenAndServe(config.App.Port, r); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
	defer PSQLConn.CloseConn()
}
