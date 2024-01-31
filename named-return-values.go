package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Thomas"
	middleName = "Averdamen"
	lastName = "Ardiansah"

	return firstName, middleName, lastName

}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}