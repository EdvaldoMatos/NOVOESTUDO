package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco endereco
}
type endereco struct {
	rua    string
	numero uint16
}

func main() {

	fmt.Println("Structs")

	var u usuario
	u.nome = "João"
	u.idade = 30
	u.endereco = endereco{rua: "Rua A", numero: 123}
	fmt.Println(u)

	u2 := usuario{nome: "Maria", idade: 25}
	fmt.Println(u2)

	u3 := usuario{idade: 40}
	fmt.Println(u3)
}
