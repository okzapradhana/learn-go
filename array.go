package main

import "fmt"

func main() {
	var movies [3]string
	movies[0] = "NKTCHI"
	movies[1] = "Dilan 1990"
	movies[2] = "Dilan 1991"

	fmt.Println("isi array", movies)
	fmt.Println(movies[0], movies[1], movies[2])

	var fruits = [3]string{"apel", "jeruk", "naga"} //inisialisasi array secara horizontal
	fmt.Println("Jumlah elemen array\t", len(fruits))
	fmt.Println("Isi semua elemen\t", fruits)

	var months = [3]string{
		"januari",
		"februari",
		"maret",
	}

	fmt.Println("Isi Array Bulan", months)
	var grades = [3]string{"A", "B", "C"}
	fmt.Println("Jumlah elemen array grade\t:", len(grades))
	fmt.Println("Isi array\t\t\t:", grades)

	/*
		multidimensional array
		- pada subdimensi, bisa menuliskan ukuran arraynya. bisa juga tidak.
	*/
	var numbers1 = [3][3]int{[3]int{1, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [3][3]int{{1, 2, 3}, {2, 3, 4}}

	fmt.Println("Numbers1", numbers1)
	fmt.Println("Numbers2", numbers2)

	var numbers = [4]uint8{1, 2, 3, 4}

	for _, number := range numbers {
		fmt.Println("Elemen angka", number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Elemen", numbers[i])
	}

	//jika butuh indeks aja, bisa pakai underscore setelah i, atau i saja.
	fmt.Println("Looping Indeks aja")
	for i := range numbers {
		fmt.Println(i)
	}

	//alokasi elemen array dengan keyword make
	var books = make([]string, 2)
	books[0] = "Apapun"
	books[1] = "Pasti Jadi"

	fmt.Println("List Buku", books)
}
