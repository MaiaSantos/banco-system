package main

import (
	"BancoSys/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDousuario1 := contas.ContaPoupanca{}
	contaDousuario1.Depositar(100)
	PagarBoleto(&contaDousuario1, 60)

	fmt.Println(contaDousuario1.ExibiSaldo())

	contaDousuario2 := contas.ContaCorrente{}
	contaDousuario2.Depositar(200)
	PagarBoleto(&contaDousuario2, 70)

	fmt.Println(contaDousuario2.ExibiSaldo())
}
