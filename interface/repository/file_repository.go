package repository

import (
	"log"
	"os"
)

// Checks and creates if file doesn't exist
func CreateFileIfNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, errCreate := os.Create(path)
		if errCreate != nil {
			log.Fatalf("failed creating file: %s", err)
			return false
		}
	}
	return true
}
