package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "bootcamp/domain/model"
	. "bootcamp/interface/repository"
)

type pokemonController struct {
	*gin.Engine
	repo PokemonRepository
}
type PokemonController interface {
	Get() gin.HandlerFunc
	GetAsync() gin.HandlerFunc
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

func (this *pokemonController) GetAsync() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var uri PokemonAsyncUriQueryParams
		if ctx.ShouldBindQuery(&uri) == nil {
			if uri.Type != "odd" && uri.Type != "even" {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Type value should be even or odd"})
			} else if uri.Items <= 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Items value should be greater than 0"})
			} else if uri.ItemsPerWorker <= 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Items per worker value should be greater than 0"})
			} else if uri.ItemsPerWorker > uri.Items {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Items value should be greater than Items per worker"})
			} else {
				pkmn, err := this.repo.GetAsync(uri)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err})
				} else {
					ctx.JSON(http.StatusOK, gin.H{"length": len(pkmn), "data": pkmn})
				}
			}

		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid or missing params, please verify"})
		}
	}
}
