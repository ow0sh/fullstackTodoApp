package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

type TypeAppConfig struct {
	Port string `json:"PORT"`
}

type TypePSQLConfig struct {
	Dsn      string `json:"dsn"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Dbname   string `json:"dbname"`
}

type TypeConfig struct {
	App  TypeAppConfig  `json:"app"`
	PSQL TypePSQLConfig `json:"psql"`
}

func InitConfig(log *logrus.Logger) (TypeConfig, error) {
	configFilename := "default.json"
	var config TypeConfig

	configFile, err := ioutil.ReadFile("./config/" + configFilename)
	if err != nil {
		log.Error("failed to open config file")
		return config, errors.Wrap(err, "failed to open config file")
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Error("failed to open config file")
		return config, errors.Wrap(err, "failed to open config file")
	}

	log.Info("config initialized")
	return config, nil
}
