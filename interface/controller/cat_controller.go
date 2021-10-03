package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "bootcamp/interface/repository"
)

type catController struct {
	*gin.Engine
}
type CatController interface {
	GetCat() gin.HandlerFunc
}

//Cat Controller instance
func NewCatController(e *gin.Engine) *catController {
	return &catController{e}
}

//Gets a cat from a external API
func (this *catController) GetCat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cat, err := NewCatRepository().GetCatFromApi()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, cat)
	}
}
