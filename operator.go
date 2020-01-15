package main

import "fmt"

func main() {
	var result = ((2 * 5) + (10 * 3))

	fmt.Println("Result", result)

	//logic operator
	var isLeft, isRight = true, false
	var bool_result_and = isLeft && isRight
	var bool_result_or = isLeft || isRight

	fmt.Printf("Right and Left? %t\n", bool_result_and)
	fmt.Printf("Right or Left? %t\n", bool_result_or)

}
