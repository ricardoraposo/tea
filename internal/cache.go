package internal

import (
	"os"
	"path/filepath"
	"runtime"
)

func darwinCacheDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cacheDir := filepath.Join(homeDir, "Library", "Caches")
	return cacheDir, err
}

func linuxCacheDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cacheDir := filepath.Join(homeDir, ".cache")
	return cacheDir, err
}

func getCacheDir() (string, error) {
	var cacheDir string
	var err error
	os := runtime.GOOS
	switch os {
	case "darwin":
		cacheDir, err = darwinCacheDir()
	default:
		cacheDir, err = linuxCacheDir()
	}
	return cacheDir, err
}

func CreateCacheDir() (string, error) {
	cacheDir, err := getCacheDir()
	if err != nil {
		return "", err
	}
	folderPath := filepath.Join(cacheDir, "tea")
	err = os.Mkdir(folderPath, 0755)
	if err != nil {
		return folderPath, err
	}
	return folderPath, nil
}
