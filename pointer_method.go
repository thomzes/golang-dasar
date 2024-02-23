package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	thomas := Man{"Thomas"}
	thomas.Married()

	fmt.Println(thomas.Name)
}