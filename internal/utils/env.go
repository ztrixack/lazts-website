package utils

import "os"

func LookupEnv(key string, _default string) string {
	ret, ok := os.LookupEnv(key)
	if !ok {
		return _default
	}

	return ret
}
