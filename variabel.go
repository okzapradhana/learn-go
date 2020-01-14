package main

import "fmt"

func main() {
	//variabel dengan tipe data
	var nama string = "okza"
	var status string
	status = "data enthusiast"

	fmt.Printf("Nama saya %s, saya seorang %s\n", nama, status)

	//variabel tanpa tipe data (type inference) dengan operand ':='
	var hobi string = "bermain basket"
	lagu_fav := "jepang"

	fmt.Printf("Hobi saya %s dengan lagu favorit %s\n", hobi, lagu_fav)

	//multi variabel
	var satu, dua, tiga string
	satu, dua, tiga = "1st", "2nd", "3rd"

	var empat, lima, enam string = "4th", "5th", "6th"

	tujuh, delapan := "7th", "8th"

	//inference multi variabel dengan tipe data berbeda
	angka_fav, makanan_fav, tinggi := 17, "ayam", 170.5

	fmt.Println(satu, dua, tiga, empat, lima, enam, tujuh, delapan, angka_fav, makanan_fav, tinggi)

	//underscore, predefined sehingga tidak membutuhkan operand :=
	_ = "Tidak bisa dilihat hasilnya"

	//multivariabel dengan underscore
	jarak, _ := 10, "Tes"

	fmt.Println("Jaraknya", jarak)

	//variabel pointer
	nama_operator_pointer := new(string)
	fmt.Println("alamat memori", nama_operator_pointer)
	fmt.Println("hasil dereference =", *nama_operator_pointer)
}
