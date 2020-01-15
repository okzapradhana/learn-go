package main

import "fmt"

func main() {
	const name string = "okza"
	const name_inference = "okza"

	fmt.Printf("Nama tanpa inference %s\n", name)
	fmt.Printf("Nama dengan inference %s\n", name_inference)

	//fmt.Print() => semua argumen dalam fungsi Print tidak dipisah oleh spasi, Berbeda seperti Println
	fmt.Print("okza pradhana\n")
	fmt.Print("okza", " ", "pradhana\n")

	fmt.Println("okza", "pradhana")

}
