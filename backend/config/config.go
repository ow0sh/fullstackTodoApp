package config

import (
	"encoding/json"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config interface {
	DB() *sqlx.DB
	Log() *logrus.Logger
}

type config struct {
	db
	logger
}

func NewConfig(configPath string) (Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config")
	}

	cfg := config{}
	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to decode config")
	}

	if err = cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "failde to validate config")
	}

	return &cfg, nil
}
