package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name

	return hello
}

func main() {
	result := getHello("Anto")
	fmt.Println(result)
}