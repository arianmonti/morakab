package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string `env:"APP_NAME" env-default:"Morakkab"`
	AppVersion  string `env:"APP_VERSION" env-default:"0.01"`
	Port        string `env:"PORT" env-default:":5000"`
	DatabaseURL string `env:"DATABASE_URL" env-default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
	JWTSecret   string `env:"JWT_SECRET" env-default:"secret"`
}

var Cfg Config

func init() {
	godotenv.Load()
	if err := cleanenv.ReadEnv(&Cfg); err != nil {
		panic(err)
	}
}
