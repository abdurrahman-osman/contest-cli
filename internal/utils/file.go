package utils

import (
	"os"
	"strings"
)

// ExpandHomeDir resolves paths starting with "~/" to the absolute user home directory.
func ExpandHomeDir(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return home + path[1:], nil
	}
	return path, nil
}
