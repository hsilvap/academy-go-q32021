package main

import (
	. "pokemon/controller"

	"github.com/gin-gonic/gin"
)

//Principal function
func main() {
	//R is the abbreviation of router
	r := gin.Default()
	NewPokemonController(r).Router()
	r.Run(":8080")
}
