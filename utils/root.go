package utils

import (
	"os"
	"path/filepath"
)

func GetProjectRoot(subDir1, subDir2 string) string {
	cwd, _ := os.Getwd()
	baseDir := FindRoot(cwd)
	return filepath.Join(baseDir, subDir1, subDir2)
}

func FindRoot(cwd string) string {
	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); os.IsNotExist(err) {
			parentDir := filepath.Dir(cwd)
			if parentDir == cwd {
				return cwd
			}
			cwd = parentDir
		} else {
			return cwd
		}
	}
}
