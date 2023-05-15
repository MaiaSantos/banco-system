package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroConta   int
	numeroAgencia int
	saldo         float64
}

func main() {
	contaDoUsuario1 := ContaCorrente{titular: "João", numeroConta: 589, numeroAgencia: 123456, saldo: 1231.12}

	contaDoUsuario2 := ContaCorrente{"Dinho", 321, 123432, 1232.02}

	fmt.Println(contaDoUsuario1)
	fmt.Println(contaDoUsuario2)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(&contaDaCris) // & para ver o endereço na memória
	fmt.Println(&contaDaCris2)

	fmt.Println(*contaDaCris == *contaDaCris2) // comparação estritamente igual
}
