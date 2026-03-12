package utils

import (
	"fmt"

	"exemple.com/bank/account"
)

func Sacar() {
	fmt.Printf("\nInforme o valor a ser sacado: ")
	var valor float64
	fmt.Scanln(&valor)

	if(valor <= 0 || valor > account.Saldo) {
		fmt.Println("\nValor inválido!")
		return
	}

	account.Saldo -= valor
	fmt.Printf("\nSaque de R$%.2f realizado com sucesso!\n", valor)
}