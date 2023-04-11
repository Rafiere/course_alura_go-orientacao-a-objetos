package main

import (
	"fmt"
	"go-orientacao-a-objetos/contas"
	"reflect"
)

func main() {

	var titular string = "Testando"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 500.50

	fmt.Println("Hello world, ", titular+".")
	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	contaDoRafa := contas.ContaCorrente{
		Titular:       "Rafa",
		NumeroAgencia: 500,
		NumeroConta:   50,
		Saldo:         500.50,
	}

	fmt.Println(contaDoRafa)
	fmt.Println(reflect.TypeOf(contaDoRafa))

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"

	/* Quando usamos o &variavel, estamos acessando o endereço dessa variável. Quando
	usamos o "*variavel", estamos acessando o valor que está nesse endereço. A sintaxe
	"*variavel" só é válida para acessarmos o valor em que o ponteiro está apontando. */

	fmt.Println(reflect.TypeOf(*contaDaCris))

	fmt.Println(contaDoRafa.Saldo)
	fmt.Println(contaDoRafa.Sacar(430))
	fmt.Println("Novo valor: R$", contaDoRafa.saldo)

	/* ... */

	contaDaSilvia := contas.ContaCorrente{Titular: "Sílvia", Saldo: 300}
	contaDaRenata := contas.ContaCorrente{Titular: "Renata", Saldo: 100}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDaRenata)

	status := contaDaRenata.Transferir(200, &contaDaSilvia)

	fmt.Println(status)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDaRenata)
}
