package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetEnvInt(key string) int {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	result, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
	return result
}
