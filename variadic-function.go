package main

import "fmt"

func sumAll(numbers ...int)int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10))

	// jika ingin menambahkan data slice ke dalam variable argumen maka tambahkan "..." di parameter function tersebut
	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))

}