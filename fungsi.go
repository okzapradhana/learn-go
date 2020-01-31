package main

import (
	"fmt"
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
