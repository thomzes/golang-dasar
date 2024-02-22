package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address) //pointer (&) bisa diganti dengan new()
	var alamat2 *Address = alamat1

	alamat2.Country = "Belanda"

	fmt.Print(alamat1)
	fmt.Print(alamat2)
}