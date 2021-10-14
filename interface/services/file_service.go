package services

import (
	"log"
	"os"
)

type fileService struct {
}

type FileService interface {
	CreateFileIfNotExists(path string) bool
	OpenCsvFile(path string) (*os.File, error)
}

//FileService instance
func NewFileService() *fileService {
	return &fileService{}
}

// Checks and creates if file doesn't exist
func (f fileService) CreateFileIfNotExists(path string) bool {
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
func (f fileService) OpenCsvFile(path string) (*os.File, error) {
	csvFile, err := os.Open(path)

	if err != nil {
		log.Fatalf("failed loading file: %s", err)
		return nil, err
	}

	return csvFile, nil
}
