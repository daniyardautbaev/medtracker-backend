package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer HTTPServerConfig `envPrefix:"HTTP_"`
	DB         DBConfig         `envPrefix:"DB_"`
}

type HTTPServerConfig struct {
	Port string `env:"PORT" envDefault:"8080"`
}

type DBConfig struct {
	URI string `env:"MONGO_URI" envDefault:"mongodb://localhost:27017"`
}

func NewConfig(filenames ...string) (*Config, error) {
	_ = godotenv.Load(filenames...)
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
