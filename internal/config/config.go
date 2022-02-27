package config

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var ErrConfigInit = errors.New("failed to init config")

// Config struct that depends configuration of App.
type Config struct {
	Server struct {
		Port         string `yaml:"port" env:"PORT" env-default:"8080"`
		WriteTimeout int    `yaml:"write_timeout" env-default:"15"`
		ReadTimeout  int    `yaml:"read_timeout" env-default:"15"`
		IdleTimeout  int    `yaml:"idle_timeout" env-default:"30"`
	}
	Postgres struct {
		Host     string `yaml:"host" env:"PG_HOST" env-default:"localhost"`
		Port     string `yaml:"port" env:"PG_PORT" env-default:"5432"`
		User     string `yaml:"user" env:"PG_USER" env-default:"postgres"`
		Password string `yaml:"password" env:"PG_PASSWORD" env-default:"postgres"`
		DBName   string `yaml:"db_name" env:"PG_NAME" env-default:"finance"`
	}
	Redis struct {
		Host     string `yaml:"host" env:"RD_HOST" env-default:"localhost"`
		Port     string `yaml:"port" env:"RD_PORT" env-default:"6379"`
		Password string `yaml:"password" env:"RD_PASSWORD" env-default:"redis"`
	}
	Logger struct {
		Path  string `yaml:"path" env:"LOG_PATH" env-default:"./logs/logs.log"`
		Level int    `yaml:"level" env:"LOG_LEVEL" env-default:"6"`
	}
}

// nolint: gochecknoglobals
var instance *Config // TODO Mustn't be singleton

// nolint: gochecknoglobals
var once sync.Once

// GetConfig return pointer to config. Config is singleton.
func GetConfig(path string) (c *Config, err error) {
	once.Do(func() {
		log.Print("reading server config file")
		instance = &Config{}
		if path == "" {
			path = "./config/config.yaml"
		}
		if err = cleanenv.ReadConfig(path, instance); err != nil {
			once = sync.Once{}

			instance = nil

			return
		}
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrConfigInit.Error(), err)
	}

	if instance == nil {
		return nil, ErrConfigInit
	}

	return instance, nil
}
