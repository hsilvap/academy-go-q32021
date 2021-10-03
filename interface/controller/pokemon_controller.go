package controller

import (
	"github.com/gin-gonic/gin"

	. "bootcamp/interface/repository"
)

type pokemonController struct {
	*gin.Engine
}
type PokemonController interface {
	GetPokemon() error
}

//New Pokemon Controller instance
func NewPokemonController(e *gin.Engine) *pokemonController {
	return &pokemonController{e}
}

//Reads pokemon from a csv file
func (this *pokemonController) GetPokemon() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data, err = GetAllPokemon()
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(200, data)
	}
}
