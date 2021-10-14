package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	. "bootcamp/infraestructure/router"
	. "bootcamp/interface/controller"
	. "bootcamp/interface/repository"
)

//Principal function
func main() {
	engine := gin.Default()
	r := SetupRouter(engine, NewCatController(engine, NewCatRepository()), NewPokemonController(engine, NewPokemonRepository()))
	if err := r.Run(":3000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
