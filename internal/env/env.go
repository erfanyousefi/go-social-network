package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return val
}
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	if valAsInt, err := strconv.Atoi(val); err != nil {
		return fallback
	} else {
		return valAsInt
	}
}
func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	if valAsBool, err := strconv.ParseBool(val); err != nil {
		return fallback
	} else {
		return valAsBool
	}
}
