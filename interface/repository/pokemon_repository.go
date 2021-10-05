package repository

import (
	"encoding/csv"
	"path/filepath"
	"strconv"

	. "bootcamp/domain/model"
	. "bootcamp/interface/services"
)

type pokemonRepository struct {
}
type PokemonRepository interface {
	GetAll() ([]Pokemon, error)
}

//Pokemon Repository instance
func NewPokemonRepository() *pokemonRepository {
	return &pokemonRepository{}
}

var (
	pkmnpath     = "infraestructure/filestore/pokemon"
	pkmnfilename = "pokemon.csv"
)

//Reads pokemons from a CSV
func (p pokemonRepository) GetAll() ([]Pokemon, error) {

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
