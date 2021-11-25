package settings

import (
	"os"
)

func GetEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return defaultValue
}

func Load() {
}
