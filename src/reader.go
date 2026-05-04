package src

import (
	"os"
	"strings"
)

// ReadBanner loads the banner file and splits it into lines
func ReadBanner(filepath string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// Normalizes Windows (\r\n) to Unix (\n) and splits by line
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(content, "\n"), nil
}
