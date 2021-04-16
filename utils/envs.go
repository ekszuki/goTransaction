package utils

import "os"

func GetEnv(env, fallback string) string {
	lookupEnv, found := os.LookupEnv(env)
	if !found {
		return fallback
	}
	return lookupEnv
}
