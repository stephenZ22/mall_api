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
	fmt.Println("Welcome to Mall")

	r := routers.RegisterRouters()

	r.Run(":3000")

}

// func test_bcrypt(password string) bool {
// 	// hashedPw, err := bcryptPw(password)
// 	hashedPw := "$2a$10$KZzAnOKOXOKtqcG4K0B/Iue2w5n0UesqqivVWqBVwMnl.ZmrH4Zwu"
// 	// if err != nil {
// 	// 	return false
// 	// }

// 	compare_err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(password))
// 	if compare_err != nil {
// 		fmt.Print("bbbbb")
// 		return false
// 	}

// 	return true
// }

// func bcryptPw(password string) (string, error) {
// 	bcryptPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println("hashed pw:", string(bcryptPw))

// 	return string(bcryptPw), nil
// }
