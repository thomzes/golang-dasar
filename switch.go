package main

import "fmt"

func main() {
	name := "Thomas"

	switch name {
	case "Thomas":
		fmt.Println("Hello Thomas")
	case "Dono":
		fmt.Println("Hai Kamu Salah Orang")
	default:
		fmt.Println("Kamu Bener Bener Salah Orang!")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama sudah benar!")
	case false:
		fmt.Println("Nama terlalu panjang!")
	}

	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang!")
	case length < 1:
		fmt.Println("Nama terlalu pendek!")
	default:
		fmt.Println("Nama sudah pas mantap")
	}

}
