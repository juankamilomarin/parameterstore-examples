package main

import (
	"fmt"
	"os"
)

type EnvVarParameterStore struct{}

func (ps EnvVarParameterStore) GetParams(paramNames []string) (map[string]string, error) {
	paramMap := map[string]string{}
	for _, key := range paramNames {
		value := os.Getenv(key)
		if value == "" {
			return map[string]string{}, fmt.Errorf(`environment variable %s does not exist`, key)
		}
		paramMap[key] = value
	}
	return paramMap, nil
}
