package repository

import (
	"log"
	"os"
	"path/filepath"
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

// Tries to open a csv file
func OpenCsvFile(path string) (*os.File, error) {
	csvFile, err := os.Open(filepath.Join(pkmnpath, filepath.Base(pkmnfilename)))

	if err != nil {
		log.Fatalf("failed loading file: %s", err)
		return nil, err
	}

	return csvFile, nil
}
