package utils

import (
	"os"
	"github.com/joho/godotenv"
)

func GetEnv(key, def string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	val := os.Getenv(key)
	if len(val) > 0 {
		return val
	}

	return def
}