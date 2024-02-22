package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Pemalang", "Jawa Tengah", "Indonesia"}
	// address2 := address1 //copy value karna tidak menggunakan pointer (&)

	// address2.City = "Bandung"
	// fmt.Println(address1) //tidak berubah
	// fmt.Println(address2) //berubah menjadi Bandung

	var address1 Address = Address{"Pemalang", "Jawa Tengah", "Indonesia"}
	var address2 *Address = &address1 //reference value dari address1, karena menggunakan pointer (&)

	address2.City = "Bandung"
	fmt.Println(address1) //ikut berubah
	fmt.Println(address2) //berubah menjadi Bandung


}