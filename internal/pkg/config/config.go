package config

import (
	"os"

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
	configFile, err := os.ReadFile("internal\\pkg\\config\\config.yaml")
	if err != nil {
		return nil, err
	}
	db := &Database{}
	server := &ServerData{}
	cfg := &Config{Db: db, Server: server}
	err = yaml.Unmarshal(configFile, cfg)

	return cfg, err
}
