package main

import (
	"fmt"
	"math"
)
func main() {
	/*go tidak mendukung ternary condition (bool var ? if var true: if var false)

	blok if else*/
	var value = 70
	if value > 70 {
		fmt.Print("Menakjubkan!\n")
	} else if value > 60 && value < 70 {
		fmt.Print("Memuaskan!\n")
	} else {
		fmt.Print("Berjuang lagi ya!\n")
	}

	var nilai float32 = 100
	var total float32 = 1000

	fmt.Printf("Math Trunc dengan argumen Float %f\n", math.Trunc(0.1))
	fmt.Print("Gunakan fungsi konversi ke float jika ingin compare pada blok if: ", float32(0.5), "\n")

	//hanya bisa menggunakan inference type pada blok if untuk definisi variabel
	if rataRata := uint32((nilai / total) * 100); float32(rataRata) > float32(0.2) {
		fmt.Printf("Good! You got Top %d%s\n", rataRata, "%")
	} else {
		fmt.Printf("Nice! You got Top %d%s\n", rataRata, "%")
	}

	//switch case, tidak lanjut ke case berikutnya walau tidak ada keyword break
	var score = 8
	switch score {
	case 10:
		fmt.Print("Perfect!\n")
	case 9:
		fmt.Print("Awesome!\n")
	case 8, 7, 6, 5: //8 atau 7 atau 6 atau 5
		fmt.Print("Great!\n")
		fmt.Print("Lets do more exercise!\n")
	default:
		{
			fmt.Print("Do more!\n")
			fmt.Print("You can do it!\n")
		}
		//menggunakan kurung kurawal akan lebih rapih dan mudah dibaca
	}

	var height = 160
	switch { //switch dengan gaya if else, setelah keyword switch tidak diikuti dengan nama variabel
	case (height > 160) && (height < 180):
		fmt.Print("Sangat tinggi!\n")
	case (height < 160 && height > 165):
		fmt.Print("Good height!\n")
	default:
		fmt.Print("Kurang tinggi\n")
	}

	//menggunakan fallthrough
	var grade = 5
	switch {
	case grade < 4:
		fmt.Print("Learn\n")
	case (grade >= 5) && (grade < 10):
		fmt.Print("Good!\n")
		fallthrough
	case (grade > 10) && (grade < 15):
		fmt.Print("Awesome!\n")
	default:
		fmt.Print("Magnificent!\n")
	}

	//nested if else
	var price = 1000
	if price > 1000 {
		if price < 2000 {
			fmt.Print("Harganya lebih dari 1000 kurang dari 2000\n")
		} else {
			fmt.Print("Harganya lebih dari 2000\n")
		}
	} else {
		switch {
		case (price > 500 && price < 700):
			fmt.Print("Harganya lebih dari 500 kurang dari 700\n")
		default:
			fmt.Print("Diluar range!\n")
		}
	}
}
