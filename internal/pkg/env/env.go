package env

import (
	"os"
)

func ReadString(key string, fallback string) string {
	env := os.Getenv(key)
	return env
}
