package main

import "fmt"

func main() {
	person := map[string]string{
		"Name": "Thomas",
		"Desc": "Suka Ngoding",
	}

	fmt.Println(person);
	fmt.Println(person["Name"]);
	fmt.Println(person["Desc"]);

	book := make(map[string]string);
	book["title"] = "Buku Go-Lang"
	book["author"] = "Thomas Ardi"
	book["wrong"] = "Oops"
	
	fmt.Println(book);

	delete(book, "wrong");

	fmt.Println(book);

}