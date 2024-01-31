package main

import "fmt"

func main() {
	names := [...]string{
		"Thomas",
		"Ardi",
		"Ansah",
		"Diki",
		"Dony",
		"Bana",
		"Rara",
	}

	slice := names[2:5]
	slice1 := names[:3]
	var slice2 []string = names[4:]
	slice3 := names[:]

	fmt.Println(slice);
	fmt.Println(len(slice)); //mengecek panjang slice

	fmt.Println(slice1);
	fmt.Println(cap(slice1)); //menfgecek kapasitas slice

	fmt.Println(slice2);
	fmt.Println(append(slice2, "Data Baru")); //menambah data baru pada slice

	fmt.Println(slice3);

	// ================================== \\

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:];
	daySlice1[0] = "Sabtu Baru";
	daySlice1[1] = "Minggu Baru";
	fmt.Println(days); //[Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru");
	daySlice2[0] = "Waduh";
	fmt.Println(daySlice2); //[Waduh, Minggu Baru, Libur Baru]
	fmt.Println(days) //[Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	// ============= membuat slice dengan make()=================\\

	newSlice := make([]string, 2, 5);
	newSlice[0] = "Thomas";
	newSlice[1] = "Ardiansah";
	// newSlice[2] = "Ardiansah"; //error harusnya kalau mau menambahkan slice menggunakan append()

	fmt.Println(newSlice);
	fmt.Println(len(newSlice));
	fmt.Println(cap(newSlice));

	newSlice2 := append(newSlice, "DataBaru");
	fmt.Println(newSlice2);
	fmt.Println(len(newSlice2));
	fmt.Println(cap(newSlice2));

	newSlice2[0] = "Jhin";
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice \\

	fromSlice := days[:];
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice);
	fmt.Println(fromSlice);
	fmt.Println(toSlice);

	// perbedaan deklarasi array dengan slice
	
	iniArray := [...]int{1,2,3,4,5,6}
	iniSlice := []int{1,2,3,4,5,6}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}