package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 90
	var lulusAbsensi bool = absensi > 90

	var lulus bool = lulusNilaiAkhir && lulusAbsensi;

	fmt.Println(lulus)
}