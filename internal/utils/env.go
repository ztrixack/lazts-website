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

func LookupUIntEnv(key string, _default uint) uint {
	ret, ok := os.LookupEnv(key)
	if !ok || ret == "" {
		return _default
	}

	if s, _ := strconv.Atoi(ret); s > 0 {
		return (uint)(s)
	}

	return _default
}
