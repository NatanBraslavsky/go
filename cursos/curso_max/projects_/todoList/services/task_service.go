package services

import (
	"fmt"

	"exemple.com/todo/models"
)

var Tasks []models.Task
var nextId = 1

func AddTask(title, description string) {
	newTask := models.Task{
		Id: nextId,
		Title: title,
		Description: description,
		Status: false,
	}
	Tasks = append(Tasks, newTask)
	nextId++
}

func ShowTasks() {
	if len(Tasks) == 0 {
		fmt.Println("Nenhuma task encontrada.")
		return
	}

	fmt.Println("\n__Tasks__")
	for _, v := range Tasks {
		status := "Pendente"
		if v.Status {
			status = "Concluída"
		}
		fmt.Printf("\n\nId: %d", v.Id)
		fmt.Printf("\nTitle: %s", v.Title)
		fmt.Printf("\nDescription: %s", v.Description)
		fmt.Printf("\nStatus: %s", status)
	}
}

func CheckTask(id int) {
	for i := range Tasks {
		if Tasks[i].Id == id {
			Tasks[i].Status = !Tasks[i].Status
			return
		}
	}
}

func DeleteTask(id int) {
	for i := range Tasks {
		if Tasks[i].Id == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			return
		}
	}
}