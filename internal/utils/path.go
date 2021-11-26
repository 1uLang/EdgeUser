package utils

import "os"

func Path() string {
	path, err := os.Getwd()
	if err == nil {
		return path
	}
	return "./"
}
