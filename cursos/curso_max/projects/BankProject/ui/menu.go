package ui

import (
	"fmt"
	"exemple.com/bank/utils"
)

func Menu() {
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
				utils.VerSaldo()
			case 2:
				utils.Depositar()
			case 3:
				utils.Sacar()
			case 4:
				fmt.Println("\nSaindo...")
				return
			default:
				fmt.Println("\nOpção inválida!")
		}
	}
}