package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Thomas"
	names[1] = "Ardi"
	names[2] = "Ansah"

	fmt.Println(names[0]);
	fmt.Println(names[1]);
	fmt.Println(names[2]);

	// deklarasi langsung
	var values = [3]int{
		90,
		22,
		100,
	}

	fmt.Println(values[2]);
	fmt.Println(len(values))

	// jika ingin menggunakan [...] maka harus ada isi dari array tersebut
	var isiData = [...]int{
		90,
		1000,
		100,
		2,
		5,
		7,
	}

	fmt.Println(isiData[1]);
	fmt.Println(len(isiData));
}