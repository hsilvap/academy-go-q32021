package main

import (
	"fmt"

	. "bootcamp/infraestructure/router"
)

//Principal function
func main() {
	r := SetupRouter()
	if err := r.Run(":3000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
