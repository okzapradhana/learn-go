package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	query, err := ioutil.ReadFile("./files/test.fql")
	if err != nil {
		log.Fatal(err)
	}

	text := string(query)
	fmt.Println(text)
}
