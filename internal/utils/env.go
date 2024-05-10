package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, def string) string {
	val := os.Getenv(key)
	if len(val) > 0 {
		return val
	}

	return def
}

func GetEnvLocal(key, def string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return GetEnv(key, def)
}
