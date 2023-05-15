package contas

import "BancoSys/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroConta, NumeroAgencia int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque >= 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado"
	} else {
		return "O Valor é incorreto ou saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64, string) {
	if valorDoDeposito >= 0 {
		c.saldo += c.saldo + valorDoDeposito
		return "Deposito de", c.saldo, "realizado"
	} else {
		return "O Valor", c.saldo, "incorreto ou depósito insuficiente"
	}

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.saldo >= valorDaTransferencia {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ExibiSaldo() float64 {
	return c.saldo
}
