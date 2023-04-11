package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

/* O "conta *ContaCorrente diz que, quando usarmos o método "sacar", ele referenciará o
objeto que chamar essa função, semelhante ao "this" do Java. */

func (conta *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= conta.Saldo

	if podeSacar {
		conta.Saldo -= valorDoSaque
		return "O saque foi realizado com sucesso."
	} else {
		return "Saldo Insuficiente!"
	}
}

func (conta *ContaCorrente) Depositar(valorDoSaque float64) (string, float64) {
	if valorDoSaque < 0 {
		return "Erro: ", conta.Saldo
	}

	conta.Saldo += valorDoSaque

	return "O depósito foi realizado com sucesso!", conta.Saldo
}

func (contaOrigem *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < contaOrigem.Saldo && valorDaTransferencia > 0 {
		contaOrigem.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
