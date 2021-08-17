package utils

import (
	"os"
)

func GetCurrentPath() (string, error) {
	return os.Executable()
}
