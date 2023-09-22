package env

import (
	"os"
	"strings"
)

func LoadEnv(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var env = make(map[string]string)

	for _, line := range strings.Split(string(content), "\n") {
		if line == "" {
			continue
		}

		before, after, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		env[before] = after
	}

	for k, v := range env {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
