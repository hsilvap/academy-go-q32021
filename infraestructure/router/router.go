package router

import (
	"github.com/gin-gonic/gin"

	. "bootcamp/interface/controller"
	. "bootcamp/interface/repository"
)

// Router setup with all available routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	getCatHandler := NewCatController(r, NewCatRepository()).Get()
	getPokemonHandler := NewPokemonController(r, NewPokemonRepository()).Get()
	getPokemonAsyncHandler := NewPokemonController(r, NewPokemonRepository()).GetAsync()

	//routes
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/cat", getCatHandler)
	pokemon := r.Group("/pokemon")
	{
		pokemon.GET("/get", getPokemonHandler)
		pokemon.GET("/get/async", getPokemonAsyncHandler)
	}
	return r
}
