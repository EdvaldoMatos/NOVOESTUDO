package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func calculosMatematicos(a int, b int) (int, int) {
	soma := a + b
	subtracao := a - b
	return soma, subtracao
}

func main() {
	resultado := soma(5, 3)
	fmt.Println(resultado)

	var f = func(txt string) string{
		fmt.Println(txt)
		return txt
	}
	
	resultadoFunc := f("Olá, mundo!")
	fmt.Println(resultadoFunc)	

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 4)
	fmt.Println("Soma:", resultadoSoma)
	fmt.Println("Subtração:", resultadoSubtracao)

	
	resultadoSoma2, _ := calculosMatematicos(10, 4) // Ignorando o resultado da subtração usando o underscore
	fmt.Println("Soma:", resultadoSoma2)

}
