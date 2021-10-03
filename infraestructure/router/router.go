package router

import (
	"github.com/gin-gonic/gin"

	. "bootcamp/interface/controller"
)

// Router setup with all available routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//routes
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/cat", NewCatController(r).GetCat())
	r.GET("/pokemon", NewPokemonController(r).GetPokemon())
	return r
}
