package main

import "fmt"

func main() {
	const firstName string = "Thomas"
	const lastName = "Ardiansah"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		namaAwal  string = "Thomass"
		namaAkhir        = "Ardiansahh"
	)

	fmt.Println(namaAwal)
	fmt.Println(namaAkhir)

	// Error
	// firstName = "Budi"
	// lastName = "Nana"

}