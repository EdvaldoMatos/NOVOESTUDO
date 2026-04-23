package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	
	err := checkmail.ValidateFormat("edvaks@gmail.com")
	if err != nil {
		fmt.Println("Email inválido")
	} else {
		fmt.Println("Email válido")
	}
}