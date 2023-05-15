package contas

import "BancoSys/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroConta, NumeroAgencia, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque >= 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado"
	} else {
		return "O Valor é incorreto ou saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64, string) {
	if valorDoDeposito >= 0 {
		c.saldo += c.saldo + valorDoDeposito
		return "Deposito de", c.saldo, "realizado"
	} else {
		return "O Valor", c.saldo, "incorreto ou depósito insuficiente"
	}

}

func (c *ContaPoupanca) ExibiSaldo() float64 {
	return c.saldo
}
