package main

import "fmt"

func main() {

	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000000000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var letra rune = 'a'
	fmt.Println(letra)

	//byte = uint8
	var b byte = 255
	fmt.Println(b)

	var numeroReal1 float32 = 3.14
	fmt.Println(numeroReal1)
	var numeroReal2 float64 = 3.14159265358979323846264338327950288419716939937510
	fmt.Println(numeroReal2)

	numeroReal3 := 3.14
	fmt.Println(numeroReal3)

	var str string = "Olá, mundo!"
	fmt.Println(str)

	var booleano bool = true
	fmt.Println(booleano)

	
	char := 'A'
	fmt.Printf("O tipo de char é: %T\n", char)
	fmt.Printf("O valor de char é: %c\n", char)

}
