package database

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		return err.Error()
	}

	return os.Getenv(key)
}
