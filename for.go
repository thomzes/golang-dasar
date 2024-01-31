package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke - ", counter);
		counter++;
	}
	fmt.Println("For Loop Selesai!")

	// dengan statement

	for angka := 1; angka <= 10; angka++ {
		fmt.Println("Perulangan angka ke - ", angka);
	}
	fmt.Println("Selesai For Loop Angka!")

	// for range manual menggunakan for loop biasa
	names := []string{
		"Thomas",
		"Ardiansah",
		"Dadang",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i]);
	}

	// for range

	for index, name := range names {
		fmt.Println("Index", index, "=", name);
	}

	// jika tidak butuh index, maka gunakan "_"
	
	for _, name := range names {
		fmt.Println(name);
	}
	
}