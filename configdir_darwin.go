package configdir

import (
	"errors"
	"os"
	"path/filepath"
)

func dir() (string, error) {
	d := os.Getenv("HOME")
	if d == "" {
		return "", errors.New("config: $HOME must be defined")
	}
	return filepath.Join(d, "Library", "Application Support"), nil
}
