package main

import "fmt"

func getFullName() (string, string) {
	return "Thomas", "Ardiansah"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// jika tidak ingin menggunakan salah satu value, bisa menggunakan "_"
	_, namaAkhir := getFullName()
	fmt.Println(namaAkhir)
}