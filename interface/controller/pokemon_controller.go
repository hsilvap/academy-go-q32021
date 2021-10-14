package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "bootcamp/interface/repository"
)

type pokemonController struct {
	*gin.Engine
	repo PokemonRepository
}
type PokemonController interface {
	Get()
}

//New Pokemon Controller instance
func NewPokemonController(e *gin.Engine, r PokemonRepository) *pokemonController {
	return &pokemonController{e, r}
}

//Reads pokemon from a csv file
func (this *pokemonController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data, err = this.repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, data)
	}
}
