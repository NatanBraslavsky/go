package main

import "fmt"

var saldo float64 = 1000

func main() {
	menu()	
}

func menu() {
	for {
		var escolha int
		fmt.Println("\n1. Ver saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Sacar")
		fmt.Println("4. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&escolha)

		switch escolha{
			case 1:
				verSaldo()
			case 2:
				depositar()
			case 3:
				sacar()
			case 4:
				fmt.Println("\nSaindo...")
				return
			default:
				fmt.Println("\nOpção inválida!")
		}
	}
}

func verSaldo() {
	fmt.Printf("\nSaldo: R$%.2f\n", saldo)
}

func depositar() {
	fmt.Printf("\nInforme o valor a ser depositado: ")
	var valor float64
	fmt.Scanln(&valor)

	if(valor <= 0) {
		fmt.Println("\nValor inválido!")
		return
	}

	saldo += valor
	fmt.Printf("\nDepósito de R$%.2f realizado com sucesso!\n", valor)
}

func sacar() {
	fmt.Printf("\nInforme o valor a ser sacado: ")
	var valor float64
	fmt.Scanln(&valor)

	if(valor <= 0 || valor > saldo) {
		fmt.Println("\nValor inválido!")
		return
	}

	saldo -= valor
	fmt.Printf("\nSaque de R$%.2f realizado com sucesso!\n", valor)
}
