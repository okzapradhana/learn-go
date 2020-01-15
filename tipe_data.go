package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		uint: bil cacah bernilai positif saja
		int: bil bulat bisa bernilai positif dan negatif
	*/
	positiveNumInference := 90
	var positiveNum uint8 = 90
	var negativeNum = -1243423644
	negativeNumInference := -1243423644

	var decimalNum = 2.62
	decimalNumInference := 2.62
	var decimalNumFloat float32 = 2.62

	var isFull bool = true

	var chatMessage string = "hi! how r u"
	var message = "hello!"
	var backtickMessage = `Nama saya "Okza", saya suka berdiskusi. 
	Happy to know u!
	`

	//int memiliki cakupan nilai yang sama antara int32 maupun int64 tergantung nilai yang diberikan
	fmt.Printf("Inference Type %s\n", reflect.TypeOf(positiveNumInference))
	fmt.Printf("Uint8 Type %s\n", reflect.TypeOf(positiveNum))
	fmt.Printf("Neg Num no Inference %s\n", reflect.TypeOf(negativeNum))
	fmt.Printf("Neg Num with Infernece %s\n", reflect.TypeOf(negativeNumInference))

	fmt.Printf("No Type Var Decimal %s\n", reflect.TypeOf(decimalNum))
	fmt.Printf("Inference Type %s\n", reflect.TypeOf(decimalNumInference))
	fmt.Printf("Float Decimal Num %s\n", reflect.TypeOf(decimalNumFloat))

	fmt.Printf("Full %t\n", isFull)

	fmt.Printf("Chat Message %s\n", chatMessage)
	fmt.Printf("Message %s\n", message)
	fmt.Printf("Backtick Message %s\n", backtickMessage)

}
