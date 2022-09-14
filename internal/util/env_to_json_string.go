package util

import (
	"fmt"
	"strings"
)

func BuildJsonStringFromEnv(data []string) (string, error) {
	jsonEnvFile := "{"
	for i, entry := range data {
		splitEntry := strings.Split(entry, "=")
		if len(splitEntry) == 1 {
			return "", fmt.Errorf("entry in env file does not have a value %s", entry)
		}
		jsonEnvFile += fmt.Sprintf("\"%s\":\"%s\"", splitEntry[0], strings.Join(splitEntry[1:], ""))
		if i != len(data)-1 {
			jsonEnvFile += ","
		} else {
			jsonEnvFile += "}"
		}
	}

	return jsonEnvFile, nil
}
