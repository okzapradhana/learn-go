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
	positiveNum_inference := 90
	var positiveNum uint8 = 90
	var negativeNum = -1243423644
	negativeNum_inference := -1243423644

	var decimalNum = 2.62
	decimalNum_inference := 2.62
	var decimalNum_float float32 = 2.62

	var is_full bool = true

	var chat_message string = "hi! how r u"
	var message = "hello!"
	var backtick_message = `Nama saya "Okza", saya suka berdiskusi. 
	Happy to know u!
	`

	//int memiliki cakupan nilai yang sama antara int32 maupun int64 tergantung nilai yang diberikan
	fmt.Printf("Inference Type %s\n", reflect.TypeOf(positiveNum_inference))
	fmt.Printf("Uint8 Type %s\n", reflect.TypeOf(positiveNum))
	fmt.Printf("Neg Num no Inference %s\n", reflect.TypeOf(negativeNum))
	fmt.Printf("Neg Num with Infernece %s\n", reflect.TypeOf(negativeNum_inference))

	fmt.Printf("No Type Var Decimal %s\n", reflect.TypeOf(decimalNum))
	fmt.Printf("Inference Type %s\n", reflect.TypeOf(decimalNum_inference))
	fmt.Printf("Float Decimal Num %s\n", reflect.TypeOf(decimalNum_float))

	fmt.Printf("Full %t\n", is_full)

	fmt.Printf("Chat Message %s\n", chat_message)
	fmt.Printf("Message %s\n", message)
	fmt.Printf("Backtick Message %s\n", backtick_message)

}
