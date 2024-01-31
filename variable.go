package main

import "fmt"

func main() {
	var name string

	name = "Thomas Ardiansah"
	fmt.Println(name);

	name = "Ardiansah Thomas"
	fmt.Println(name);

	/* deklare variable langsung berupa value 

	var name = "Thomas Ardiansah"
	fmt.Println(name);

	*/

	/* membuat variable menggunakan ":="
	note, menggunakan ":=" hanya untuk awal saja

	name := "Thomas Ardiansah"
	fmt.Println(name);

	name = "Ardiansah Thomas"
	fmt.Println(name);
	*/

	var (
		firstName = "Thomas"
		middleName = "Ardia"
		lastName = "Ansah"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)



	

}