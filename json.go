package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Type   string `json: "Type"`
	Place  string
	HasLeg bool
}

func main() {
	var jsonString = `{"Type": "carnivore", "Place": "water", "HasLeg": false}`
	var arrayJsonString = `[
		{"Type": "herbivore", "Place": "land", "HasLeg": true},
		{"Type": "omnivore", "Place": "water", "HasLeg": false}	
	]`
	var jsonData = []byte(jsonString)

	var animal Animal
	var animals []Animal
	var mapAnimal map[string]interface{}
	var intfAnimal interface{}

	//Unmarshal([]byte, interface) => untuk decode dari JSON string ke objek
	if err := json.Unmarshal([]byte(arrayJsonString), &animals); err != nil {
		fmt.Println(err.Error())
		return
	}

	json.Unmarshal(jsonData, &intfAnimal)
	var decodedAnimal = intfAnimal.(map[string]interface{}) //casting type ke map

	if err := json.Unmarshal(jsonData, &animal); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := json.Unmarshal(jsonData, &mapAnimal); err != nil {
		fmt.Println(err.Error())
		return
	}

	//Marshal(interface) => encode dari objek ke JSON String
	var newAnimals = []Animal{{"herbivore", "land", true}, {"carnivore", "land", true}}
	var objectAnimal, _ = json.Marshal(newAnimals)

	fmt.Println(animal)
	fmt.Println(animals)
	fmt.Println(animals[0].Type)
	fmt.Println(animals[1].Type)
	fmt.Println(mapAnimal)
	fmt.Println("Marshal animal", objectAnimal)
	fmt.Println("String marshal animal", string(objectAnimal))
	fmt.Println(decodedAnimal)
	fmt.Println(mapAnimal["Type"])
	fmt.Println(mapAnimal["Place"])
	fmt.Println("Type ", animal.Type)
	fmt.Println("Place ", animal.Place)
}
