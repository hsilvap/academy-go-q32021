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
	Get()
	GetAsync()
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
		var uri ConcurrentUriQueryParams
		if ctx.ShouldBindQuery(&uri) == nil {
			if uri.Type != "odd" && uri.Type != "even" {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Type value should be even or odd"})
			} else if uri.Items <= 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Items value should be greater than 0"})
			} else if uri.ItemsPerWorker <= 0 {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Items per worker value should be greater than 0"})
			} else {
				var data = this.repo.GetAsync(uri)
				ctx.JSON(http.StatusOK, gin.H{"length": len(data), "data": data})
			}

		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid or missing params, please verify"})
		}
	}
}
