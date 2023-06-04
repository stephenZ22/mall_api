package main

import (
	"MallApi/initializers"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.DataBaseInitMigrate()
}

func main() {
	fmt.Println("Hello Mall")
}
