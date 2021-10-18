package router

import (
	"github.com/gin-gonic/gin"

	. "bootcamp/interface/controller"
)

// Router setup with all available routes
func SetupRouter(r *gin.Engine, catController CatController, pokemonController PokemonController) *gin.Engine {

	//routes
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/cat", catController.Get())
	pokemon := r.Group("/pokemon")
	{
		pokemon.GET("/get", pokemonController.Get())
		pokemon.GET("/get/async", pokemonController.GetAsync())
	}
	return r
}
