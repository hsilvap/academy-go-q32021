package router

import (
	"github.com/gin-gonic/gin"

	. "bootcamp/interface/controller"
	. "bootcamp/interface/repository"
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
	r.GET("/cat", NewCatController(r, NewCatRepository()).Get())
	r.GET("/pokemon", NewPokemonController(r, NewPokemonRepository()).Get())
	return r
}
