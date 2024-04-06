package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Database struct {
	Hostname string `yaml:"hostname"`
	Database string `yaml:"database"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ServerData struct {
	Port string `yaml:"port"`
}

type Config struct {
	Db     *Database   `yaml:"db"`
	Server *ServerData `yaml:"server"`
}

func Configure() (*Config, error) {
	f, err := os.Open("C:\\Users\\jbob0\\source\\repos\\Test\\internal\\pkg\\config\\config.yaml")

	cfg := &Config{}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func ConfigureFromEnv() (*Config, error) {
	db := &Database{}
	server := &ServerData{}

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	if os.Getenv("DB_HOSTNAME") == "" {
		return nil, errors.New(" DB_HOSTNAME property not found")
	}
	db.Hostname = os.Getenv("DB_HOSTNAME")

	if os.Getenv("DB_DATABASE") == "" {
		return nil, errors.New(" DB_DATABASE property not found")
	}
	db.Database = os.Getenv("DB_DATABASE")

	if os.Getenv("DB_PORT") == "" {
		return nil, errors.New(" DB_PORT property not found")
	}
	db.Port = os.Getenv("DB_PORT")

	db.Username = os.Getenv("DB_USERNAME")
	db.Password = os.Getenv("DB_PASSWORD")

	if os.Getenv("SERVER_PORT") == "" {
		return nil, errors.New(" SERVER_PORT property not found")
	}
	server.Port = os.Getenv("SERVER_PORT")

	return &Config{Db: db, Server: server}, nil
}
