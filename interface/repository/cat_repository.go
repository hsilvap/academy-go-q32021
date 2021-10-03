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
)

const apikey = "36f66790-4986-4490-a054-b26952d733fc"
const url = "https://api.thecatapi.com/v1/images/search"
const path = "infraestructure/filestore/cat"
const filename = "cats.csv"

//Reads cat data from a external web API
func GetCatFromApi() ([]Cat, error) {
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
	writeCatData(catDto)
	return catDto, nil
}

func writeCatData(cats []Cat) {
	CreateFileIfNotExists(filepath.Join(path, filepath.Base(filename)))
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
