package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var thomas Customer
	thomas.Name = "Thomas Ardiansah"
	thomas.Address = "Jakarta"
	thomas.Age = 11

	fmt.Println(thomas)
	fmt.Println(thomas.Name)
	fmt.Println(thomas.Address)
	fmt.Println(thomas.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     22,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Malang", 100}
	fmt.Println(budi)

	// karena function sayHello sudah menempel pada struct, maka harus ada object dari struct yang dipanggil, contoh disini adalah budi dan thomas
	budi.sayHello("Bana")
	thomas.sayHello("Dadang")

}
