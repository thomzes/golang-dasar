package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Pemalang", "Jawa Tengah", "Indonesia"}
	var address2 *Address = &address1 //reference value dari address1, karena menggunakan pointer (&)

	address2.City = "Bandung"
	fmt.Println(address1) //ikut berubah
	fmt.Println(address2) //berubah menjadi Bandung

	// address2 = &Address{"Slipi", "Jakarta Barat", "Medan"} //data tidak reference dari addres1, karna menggunakan reference dari &Address
	*address2 = Address{"Slipi", "Jakarta Barat", "Medan"} //ketika menggunakan operator asteris (*) maka data yang bersangkutan langsung juga akan berubah
	fmt.Println(address1)
	fmt.Println(address2)

}