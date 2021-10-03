package repository

import (
	. "bootcamp/domain/model"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const pkmnpath = "infraestructure/filestore/pokemon"
const pkmnfilename = "pokemon.csv"

//Reads pokemons from a CSV
func GetAllPokemon() ([]Pokemon, error) {
	csvFile, err := os.Open(filepath.Join(pkmnpath, filepath.Base(pkmnfilename)))

	if err != nil {
		log.Fatalf("failed loading file: %s", err)
		return nil, err
	}

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()

	var result = make([]Pokemon, 0, 0)

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		fmt.Println(id, "xxx")
		pkmn := Pokemon{
			Id:   id,
			Name: record[1],
		}
		result = append(result, pkmn)
	}
	return result, nil
}
