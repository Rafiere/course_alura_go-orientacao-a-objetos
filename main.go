package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
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
}
