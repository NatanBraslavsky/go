package utils

import (
	"fmt"

	"exemple.com/bank/account"
)

func VerSaldo() {
	fmt.Printf("\nSaldo: R$%.2f\n", account.Saldo)
}