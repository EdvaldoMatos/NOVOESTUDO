package main

import "fmt"

func main() {
	var variavel1 string = "variável 1"
	variavel2 := "variável 2"

	var (
		variavel3 string = "variável 3"
		variavel4 int    = 100
	)

	variavel5, variavel6 := "variável 5", "variável 6"

	const constante1 string = "constante 1"

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(constante1)
}
