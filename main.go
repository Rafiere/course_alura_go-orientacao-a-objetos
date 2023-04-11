package main

import (
	"fmt"
	"reflect"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

/* O "conta *ContaCorrente diz que, quando usarmos o método "sacar", ele referenciará o
objeto que chamar essa função, semelhante ao "this" do Java. */

func (conta *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= conta.saldo

	if podeSacar {
		conta.saldo -= valorDoSaque
		return "O saque foi realizado com sucesso."
	} else {
		return "Saldo Insuficiente!"
	}
}

func main() {

	var titular string = "Testando"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 500.50

	fmt.Println("Hello world, ", titular+".")
	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	contaDoRafa := ContaCorrente{
		titular:       "Rafa",
		numeroAgencia: 500,
		numeroConta:   50,
		saldo:         500.50,
	}

	fmt.Println(contaDoRafa)
	fmt.Println(reflect.TypeOf(contaDoRafa))

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	/* Quando usamos o &variavel, estamos acessando o endereço dessa variável. Quando
	usamos o "*variavel", estamos acessando o valor que está nesse endereço. A sintaxe
	"*variavel" só é válida para acessarmos o valor em que o ponteiro está apontando. */

	fmt.Println(reflect.TypeOf(*contaDaCris))

	fmt.Println(contaDoRafa.saldo)
	fmt.Println(contaDoRafa.sacar(430))
	fmt.Println("Novo valor: R$", contaDoRafa.saldo)
}
