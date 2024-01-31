package main

import "fmt"

func main() {
	name := "Jokoooooooooooo"

	if name == "Thomas" {
		fmt.Println("Hello Thomas")
	} else if name == "Joko" {
		fmt.Println("Kamu Benar Joko!")
	} else {
		fmt.Println("Hi, Kamu salah!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang!")
	} else {
		fmt.Println("Nama sudah benar!")
	}

}
