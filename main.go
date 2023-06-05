package main

import (
	"MallApi/initializers"
	"MallApi/routers"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.DataBaseInitMigrate()
}

func main() {
	fmt.Println("Hello Mall")
	r := routers.RegisterRouters()

	r.Run(":3000")
}
