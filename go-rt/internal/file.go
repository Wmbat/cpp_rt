package internal

import (
	"errors"
	"os"
)

func DoesFileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}
