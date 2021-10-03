package controller

import (
	"github.com/gin-gonic/gin"

	. "bootcamp/interface/repository"
)

type catController struct {
	*gin.Engine
}
type CatController interface {
	GetCat() error
}

//Cat Controller instance
func NewCatController(e *gin.Engine) *catController {
	return &catController{e}
}

//Gets a cat from a external API
func (this *catController) GetCat() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cat, err = GetCatFromApi()
		if err != nil {
			ctx.JSON(500, err)
		}
		ctx.JSON(200, cat)
	}
}
