package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if(err != nil) {
		fmt.Println("Error at loading .env file")
	}

	return os.Getenv(key)
}

