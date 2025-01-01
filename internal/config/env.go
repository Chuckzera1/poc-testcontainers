package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_HOST     string
	DATABASE_PORT     string
	DATABASE_NAME     string
	PORT              string
}

func (e *Env) LoadEnv(envPath string) error {
	if err := godotenv.Load(envPath); err != nil {
		return fmt.Errorf("erro ao carregar o arquivo .env: %w", err)
	}

	requiredVars := []string{
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"DATABASE_HOST",
		"DATABASE_PORT",
		"DATABASE_NAME",
		"PORT",
	}

	if err := e.verifyValues(requiredVars); err != nil {
		return err
	}

	e.DATABASE_USER = os.Getenv("DATABASE_USER")
	e.DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
	e.DATABASE_HOST = os.Getenv("DATABASE_HOST")
	e.DATABASE_PORT = os.Getenv("DATABASE_PORT")
	e.DATABASE_NAME = os.Getenv("DATABASE_NAME")
	e.PORT = os.Getenv("PORT")

	return nil
}

func (e *Env) verifyValues(requiredVars []string) error {
	for _, key := range requiredVars {
		value := os.Getenv(key)
		if value == "" {
			return fmt.Errorf("a variável de ambiente '%s' não está definida ou está vazia", key)
		}
	}
	return nil
}
