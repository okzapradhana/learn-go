package main

import "fmt"

func main() {
	var month map[string]int
	month = map[string]int{} //default nilai map adalah nil

	month["januari"] = 31
	month["februari"] = 30
	fmt.Println("Januari", month["januari"])

	var person = map[string]int{
		"veronika": 11,
		"teronika": 12,
		"jerAx":    13,
	}
	var aPerson = map[string]int{}
	var bPerson = make(map[string]int)
	var cPerson = *new(map[string]int)
	fmt.Println(aPerson, bPerson, cPerson)

	fmt.Println("Person: ", person)

	//for loop di map
	for key, val := range person {
		fmt.Println(key, ": ", val)
	}

	//menghapus item pada map dengan delete
	fmt.Println(len(person))
	fmt.Println(person)

	delete(person, "veronika")

	fmt.Println("---After delete---")
	fmt.Println(len(person))
	fmt.Println(person)

	var value, isExist = person["veronika"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("not exist")
	}

	//slice of maps
	var menus = []map[string]string{
		map[string]string{"name": "ayam goreng"},
		map[string]string{"name": "ayam bakar"},
		map[string]string{"name": "sate ayam"},
	}
	fmt.Println("----Menu----")
	for idx, value := range menus {
		fmt.Println(idx, value["name"])
	}

	//versi go terbaru
	var aMenus = []map[string]string{
		{"name": "ayam goreng"},
		{"name": "ayam bakar"},
		{"name": "sate ayam"},
	}
	fmt.Println("----a Menu----")
	for idx, value := range aMenus {
		fmt.Println(idx, value["name"])
	}

}
