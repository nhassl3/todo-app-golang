package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HttpServer `yaml:"http_server"`
	DBSettings `yaml:"db"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DBSettings struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5436"`
	Username string `yaml:"username" env-default:"postgres"`
	Password string
	DBName   string `yaml:"dbname" env-default:"postgres"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("[CONFIG] CONFIG_PATH environment variable not set up")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("[CONFIG] CONFIG_PATH does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("[CONFIG] Error reading config: %s", err)
	}

	// Set up password for database
	cfg.DBSettings.Password = os.Getenv("DB_PASSWORD")
	if cfg.DBSettings.Password == "" {
		log.Fatal("[CONFIG] DB_PASSWORD environment variable not set up")
	}

	return &cfg
}
