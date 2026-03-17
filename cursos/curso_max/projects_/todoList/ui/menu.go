package ui

import (
	"fmt"

	"exemple.com/todo/services"
	
)

func Menu() {
	var escolha int
	for {
		fmt.Println("\n\n__TodoList__")
		fmt.Println("1- Adicionar tarefa")
		fmt.Println("2- Listar tarefas")
		fmt.Println("3- Concluir tarefa")
		fmt.Println("4- Remover tarefa")
		fmt.Println("5- Sair")

		fmt.Printf("Escolha: ")
		fmt.Scanln(&escolha)

		switch escolha{
		case 1:
			addTask()
		case 2:
			services.ShowTasks()
		case 3:
			checkTask()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Saindo...")
			return
		default:
			println("Opção inválida.")
		}
	}
}

func addTask() {
	var title string
	var description string

	fmt.Print("\nTítulo: ")
	fmt.Scan(&title)

	fmt.Print("Descrição: ")
	fmt.Scan(&description)

	services.AddTask(title, description)
}

func checkTask() {
	var id int
	fmt.Print("Id: ")
	fmt.Scan(&id)

	services.CheckTask(id)
}

func deleteTask() {
	var id int

	fmt.Print("Id: ")
	fmt.Scan(&id)

	services.DeleteTask(id)
}