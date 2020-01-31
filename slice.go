package main

import (
	"fmt"
	"reflect"
)

func main() {
	var sliceAnimals = []string{"kucing", "musang", "ayam", "sapi"}
	var arrayAnimals = [4]string{"kucing", "musang", "ayam", "sapi"}
	var arrayAlsoAnimals = [...]string{"kucing", "musang", "ayam", "sapi"}
	var slice = sliceAnimals[0:2]
	var aSlice = arrayAnimals[0:2]
	fmt.Println(reflect.TypeOf(sliceAnimals))
	fmt.Println(reflect.TypeOf(arrayAnimals))
	fmt.Println(reflect.TypeOf(arrayAlsoAnimals))
	fmt.Println(slice)
	fmt.Println(aSlice)

	var aAnimals = sliceAnimals[0:3]
	var bAnimals = sliceAnimals[1:]

	aAnimals[2] = "singa" //tipe data reference, mengubah nilai indeks juga akan merubah slice aslinya.
	fmt.Println("Slice Animals: ", sliceAnimals)
	fmt.Println("aAnimals: ", aAnimals)
	fmt.Println("bAnimals: ", bAnimals)
	fmt.Println(len(sliceAnimals))

	//cap -> kapasitas maksimum slice
	fmt.Println("Cap aAnimals: ", cap(aAnimals))
	fmt.Println("Cap bAnimals: ", cap(bAnimals))

	//cap pada bAnimals menjadi 3 karena slicing index dimulai dari index > 0 sehingga membuat posisi indeks 0 berubah

	var fruits = []string{"pisang", "apel", "mangga", "pepaya"}
	var aFruits = fruits[0:2] //pisang, apel -> len dari slice lebih kecil dari cap dari slice
	var bFruits = append(aFruits, "jambu")
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits) //nilai dari mangga berubah menjadi jambu. karena len(aFruits) < cap(aFruits)

	//copy -> melakukan copy slice
	dst := make([]string, 3)
	src := []string{"jan", "feb", "mar", "apr"}
	cpy := copy(dst, src) //len(dst) < len(src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(cpy) //jumlah elemen yang berhasil dicopy

	var bSrc = []string{"jan", "feb"}
	bCpy := copy(dst, bSrc)
	fmt.Println("bCpy: ", bCpy)

	//3 index slicing -> sekaligus menentukan kapasitas
	var collections = []string{"meyer", "ajer", "kilur"}
	var aColl = collections[0:2:3]
	var bColl = collections[0:1:3]
	fmt.Println("len aColl: ", len(aColl))
	fmt.Println("cap aColl: ", cap(aColl))
	fmt.Println("len bColl: ", len(bColl))
	fmt.Println("cap bColl: ", cap(bColl))
}
