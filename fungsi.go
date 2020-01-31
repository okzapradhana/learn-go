package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var books = []string{
		"Impro: Improvisation and the Theatre",
		"Ignore Everybody",
		"The Creative Habit: Learn It and Use It For Life",
		"Born Standing Up",
	}

	printMessage("Koleksi Buku", books)

	rand.Seed(time.Now().Unix())
	var randomVal int

	randomVal = randomWithRange(2, 10)
	fmt.Println("random num", randomVal)

	randomVal = randomWithRange(2, 10)
	fmt.Println("random num", randomVal)

	randomVal = randomWithRange(2, 10)
	fmt.Println("random num", randomVal)

	var area, circumference = calculateCircle(5.5)
	var pArea, pCircumference = calculatePredefinedCircle(5.5)
	fmt.Println("Circle Area", area)
	fmt.Println("Circle Circumference", circumference)

	fmt.Println("Circle Predefined Area", pArea)
	fmt.Println("Circle Predefined Circumference", pCircumference)
}

func printMessage(messange string, books []string) {
	fmt.Println(messange)
	for idx, book := range books {
		fmt.Println(idx, book)
	}

	fmt.Println("dengan Join Strings")
	fmt.Println(strings.Join(books, " "))
}

func randomWithRange(min, max int) int {
	var val = rand.Int()%(max-min+1) + min
	return val
}

//multiple return
func calculateCircle(radius float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(radius, 2)
	var circumference float64 = math.Pi * radius * 2

	return area, circumference
}

//predefined multiple return
func calculatePredefinedCircle(radius float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(radius, 2)
	circumference = math.Pi * radius * 2

	return //mengembalikan variabel area dan circumference
}
