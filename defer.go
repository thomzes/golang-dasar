package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil logging")
}

func runApplication() {
	defer logging()
	
	fmt.Println("Run apllication")
}

func main() {
	runApplication()
}