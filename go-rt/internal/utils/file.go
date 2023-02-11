package utils

import (
	"errors"
	"os"
)

func DoesFileExist(filepath string) bool {
	_, error := os.Stat(filepath)

	return !errors.Is(error, os.ErrNotExist)
}
