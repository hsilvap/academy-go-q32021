package repository

import (
	"encoding/csv"
	"path/filepath"
	"strconv"

	. "bootcamp/domain/model"
	. "bootcamp/interface/services"
)

const pkmnpath = "infraestructure/filestore/pokemon"
const pkmnfilename = "pokemon.csv"

//Reads pokemons from a CSV
func GetAllPokemon() ([]Pokemon, error) {

	var csvFile, err = NewFileService().OpenCsvFile(filepath.Join(pkmnpath, filepath.Base(pkmnfilename)))
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()

	var result = make([]Pokemon, 0, 0)

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		pkmn := Pokemon{
			Id:   id,
			Name: record[1],
		}
		result = append(result, pkmn)
	}
	return result, nil
}
