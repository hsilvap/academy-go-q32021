package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	. "bootcamp/domain/model"
	. "bootcamp/interface/services"
)

type catRepository struct {
}
type CatRepository interface {
	GetFromApi() ([]Cat, error)
	WriteData([]Cat)
}

//Cat Repository instance
func NewCatRepository() *catRepository {
	return &catRepository{}
}

var (
	apikey   = "36f66790-4986-4490-a054-b26952d733fc"
	url      = "https://api.thecatapi.com/v1/images/search"
	path     = "infraestructure/filestore/cat"
	filename = "cats.csv"
)

//Reads cat data from a external web API
func (c catRepository) GetFromApi() ([]Cat, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-api-key", apikey)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error reading cat")
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var catDto []Cat
	json.Unmarshal(bodyBytes, &catDto)
	c.WriteData(catDto)
	return catDto, nil
}

//writes cat data to a csv file
func (c catRepository) WriteData(cats []Cat) {
	NewFileService().CreateFileIfNotExists(filepath.Join(path, filepath.Base(filename)))
	csvFile, err := os.OpenFile(filepath.Join(path, filepath.Base(filename)), os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatalf("failed loading file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, row := range cats {
		row := []string{row.Id, row.Url, strconv.Itoa(row.Height), strconv.Itoa(row.Width)}
		if err := csvwriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	csvwriter.Flush()
	csvFile.Close()
}
