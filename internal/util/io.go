package util

import (
	"os"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	text := string(f)
	return strings.Split(strings.TrimSpace(text), "\n"), nil
}
