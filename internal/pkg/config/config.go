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
