package env

import (
	"fmt"
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
		
		split := strings.Split(line, "=")
		if len(split) != 2 {
			return fmt.Errorf("invalid line at: %v", line)
		}
		
		key := split[0]
		value := split[1]
		
		env[key] = value
	}
	
	for k, v := range env {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
	
	return nil
}