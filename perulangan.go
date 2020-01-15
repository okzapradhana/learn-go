package main

import "fmt"

func main() {
	//for pada Go telah mewakili semua jenis perulangan => for , foreach , while pada bahasa lain
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var a = 0
	for a < 5 {
		fmt.Println(a)
		a++
	}

	//for true
	var b = 1
	for {
		fmt.Println("Nilai saat ini", b)
		b++
		if b == 10 {
			break
		}
	}

	//break dan continue
	for i := 0; i < 10; i++ {
		if i%1 == 0 {
			fmt.Println("Ganjil")
			continue
		}

		if i%2 == 0 {
			fmt.Println("Genap")
			break
		}
	}

	//nested for
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	//menggunakan label untuk mengatur loop yang diakses saat menemui break maupun continue
	fmt.Print("Menggunakan Label pada For")
loopingTerluar:
	for i := 0; i < 6; i++ {
	loopingDidalam:
		for j := 0; j < 6; j++ {
			if i == 4 {
				fmt.Println(i, "diskip")
				break loopingTerluar
			}
			if j == 4 {
				break loopingDidalam
			}
			fmt.Printf("Matriks [%d][%d]\n", i, j)
		}
	}
}
