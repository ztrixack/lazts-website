package utils

import (
	"os"
	"strconv"
)

func LookupEnv(key string, _default string) string {
	ret, ok := os.LookupEnv(key)
	if !ok {
		return _default
	}

	return ret
}

func LookupNumericEnv(key string, _default string) string {
	ret, ok := os.LookupEnv(key)
	if !ok || ret == "" {
		return _default
	}

	if _, err := strconv.Atoi(ret); err != nil {
		return _default
	}

	return ret
}
