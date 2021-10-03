package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "bootcamp/interface/repository"
)

type pokemonController struct {
	*gin.Engine
}
type PokemonController interface {
	GetPokemon()
}

//New Pokemon Controller instance
func NewPokemonController(e *gin.Engine) *pokemonController {
	return &pokemonController{e}
}

//Reads pokemon from a csv file
func (this *pokemonController) GetPokemon() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data, err = NewPokemonRepository().GetAllPokemon()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, data)
	}
}
