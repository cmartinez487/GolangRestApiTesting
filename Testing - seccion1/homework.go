package main

import "fmt"

func main() {

	var num1 float32 = 12.4
	var num2 int8 = 6
	var num3 int8 = 4
	var num4 int8 = 2

	var restResult = num1 - float32(num2)
	var multResult = restResult * float32(num3)
	var divResult = restResult / float32(num4)

	fmt.Println("Operación <resta> realizada", restResult)

	fmt.Println("Operación <multiplicacion> realizada", multResult)

	fmt.Println("Operación <division> realizada", divResult)
}
