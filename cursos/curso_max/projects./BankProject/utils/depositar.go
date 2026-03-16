package utils

import (
	"fmt"

	"exemple.com/bank/account"
)

func Depositar() {
	fmt.Printf("\nInforme o valor a ser depositado: ")
	var valor float64
	fmt.Scanln(&valor)

	if(valor <= 0) {
		fmt.Println("\nValor inválido!")
		return
	}

	account.Saldo += valor
	fmt.Printf("\nDepósito de R$%.2f realizado com sucesso!\n", valor)
}