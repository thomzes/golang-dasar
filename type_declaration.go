package main

import "fmt"

func main() {
	type NoKTP string

	var ktpThomas NoKTP = "1111111111111"

	var contoh string = "22222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpThomas)
	fmt.Println(contohKtp)

}