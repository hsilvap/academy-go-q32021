package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "bootcamp/interface/repository"
)

type catController struct {
	*gin.Engine
	repo CatRepository
}
type CatController interface {
	Get() gin.HandlerFunc
}

//Cat Controller instance
func NewCatController(e *gin.Engine, r CatRepository) *catController {
	return &catController{e, r}
}

//Gets a cat from a external API
func (this *catController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cat, err := this.repo.GetFromApi()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, cat)
	}
}
