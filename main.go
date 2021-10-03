package main

import (
	"fmt"

	. "bootcamp/infraestructure/router"
)

//Principal function
func main() {
	r := SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
