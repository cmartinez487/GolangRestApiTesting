package main

//importacion de un solo paquete
//import "fmt"
import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Hola usuario....")
	fmt.Println("espero tengas buen dia")

	//variables y formas de asignacion de valores
	var numberTest int = 10
	var floatTest float32 = 10.1
	var stringTest string = "Variable string"
	stringTest2 := "hello"

	fmt.Println(numberTest, floatTest, stringTest, stringTest2)
	fmt.Println(reflect.TypeOf(floatTest))

}
