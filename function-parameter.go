package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hallo Nama Saya", firstName, lastName);
}

func main() {
	sayHello("Dadang", "Suratman");
}