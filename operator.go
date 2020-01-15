package main

import "fmt"

func main() {
	var result = ((2 * 5) + (10 * 3))

	fmt.Println("Result", result)

	//logic operator
	var isLeft, isRight = true, false
	var boolResultAnd = isLeft && isRight
	var boolResultOr = isLeft || isRight

	fmt.Printf("Right and Left? %t\n", boolResultAnd)
	fmt.Printf("Right or Left? %t\n", boolResultOr)

}
