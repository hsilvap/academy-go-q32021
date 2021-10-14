package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	. "bootcamp/domain/model"
	. "bootcamp/interface/services"
)

type pokemonRepository struct {
}
type PokemonRepository interface {
	GetAll() ([]Pokemon, error)
	GetAsync(params PokemonAsyncUriQueryParams) ([]*Pokemon, error)
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
		pokeInstance := Pokemon{}
		result = append(result, *pokeInstance.ParseToPokemon(record))
	}
	return result, nil
}

//Reads pokemon using a worker pool
func (p pokemonRepository) GetAsync(params PokemonAsyncUriQueryParams) ([]*Pokemon, error) {
	f, _ := os.Open(filepath.Join(pkmnpath, filepath.Base(pkmnfilename)))
	fcsv := csv.NewReader(f)
	rs := make([]*Pokemon, 0)
	numWps := params.Items / params.ItemsPerWorker
	jobs := make(chan []string, numWps)
	res := make(chan *Pokemon)
	var wg sync.WaitGroup

	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, res, params)
		}()
	}

	go func() {
		for {
			rStr, err := fcsv.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		rs = append(rs, r)
	}

	return rs, nil
}

func worker(jobs <-chan []string, results chan<- *Pokemon, params PokemonAsyncUriQueryParams) {
	totalLines := 0
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			if totalLines == params.ItemsPerWorker {
				break
			}
			id, _ := strconv.ParseUint(job[0], 10, 32)
			pokeInstance := Pokemon{}
			if params.Type == "even" && id%2 == 0 {
				results <- pokeInstance.ParseToPokemon(job)
				totalLines++
			} else if params.Type == "odd" && id%2 == 1 {
				results <- pokeInstance.ParseToPokemon(job)
				totalLines++
			}
		}
	}
}
