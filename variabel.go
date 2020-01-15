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
	laguFav := "jepang"

	fmt.Printf("Hobi saya %s dengan lagu favorit %s\n", hobi, laguFav)

	//multi variabel
	var satu, dua, tiga string
	satu, dua, tiga = "1st", "2nd", "3rd"

	var empat, lima, enam string = "4th", "5th", "6th"

	tujuh, delapan := "7th", "8th"

	//inference multi variabel dengan tipe data berbeda
	angkaFav, makananFav, tinggi := 17, "ayam", 170.5

	fmt.Println(satu, dua, tiga, empat, lima, enam, tujuh, delapan, angkaFav, makananFav, tinggi)

	//underscore, predefined sehingga tidak membutuhkan operand :=
	_ = "Tidak bisa dilihat hasilnya"

	//multivariabel dengan underscore
	jarak, _ := 10, "Tes"

	fmt.Println("Jaraknya", jarak)

	//variabel pointer
	namaOperatorPointer := new(string)
	fmt.Println("alamat memori", namaOperatorPointer)
	fmt.Println("hasil dereference =", *namaOperatorPointer)

	/*
		jumlah := 10
		jumlah = "banyak"
		fmt.Println("jumlah dengan inference", jumlah)
	*/
}
