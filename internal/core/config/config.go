package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env                string `yaml:"env" env-default:"development"`
	PGConnectionString string `yaml:"pg_connection_string"`
	HTTPServer         `yaml:"http_server"`
	Auth               `yaml:"auth"`
}

type HTTPServer struct {
	Port        string        `yaml:"port" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Auth struct {
	AccessTokenSecret  string        `yaml:"access_token_secret"`
	RefreshTokenSecret string        `yaml:"refresh_token_secret"`
	AccessTokenTTL     time.Duration `yaml:"access_token_ttl"`
	RefreshTokenTTL    time.Duration `yaml:"refresh_token_ttl"`
}

func New() (*Config, error) {
	cfgPath := "config/local.yml"
	if _, err := os.Stat(cfgPath); err != nil {
		return nil, err
	}

	var cfg Config

	err := cleanenv.ReadConfig(cfgPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
