package settings

import (
	"os"
)

func FindEnvOrDefault(key, defaultValue string) string {
	if value, defined := os.LookupEnv(key); defined {
		return value
	}
	return defaultValue
}
