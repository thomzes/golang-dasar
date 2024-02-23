package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	result := helper.SayHello("Thomas")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //tidak bisa diakses
	// fmt.Println(helper.sayGoodbye("Thomas")) //tidak bisa diakses
}