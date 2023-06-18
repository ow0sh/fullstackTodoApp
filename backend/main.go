package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/ow0sh/fullstackgo/middlware"
)

func main() {
	log := logrus.New()

	r := chi.NewRouter()
	r.Use(middlware.Logger(log))
}
