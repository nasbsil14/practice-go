package main

import (
	"fmt"
	"../web"
)

var sTest string = "123"
func main() {
	fmt.Println("Hello World!")
	s := "test"
	fmt.Println(s + sTest)

	web.Server()
}
