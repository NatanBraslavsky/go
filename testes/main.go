package main

import "fmt"

func main() {
    var valor int

    notas := []int{100, 50, 20, 10, 5, 2, 1}
    resultado := make(map[int]int)

    fmt.Print("Digite o valor: ")
    fmt.Scanf("%d", &valor)

    for _, nota := range notas {
        if valor >= nota {
            qtd := valor / nota
            resultado[nota] = qtd
            valor = valor % nota
        }
    }

    for _, nota := range notas {
        if resultado[nota] > 0 {
            fmt.Printf("%d: %d\n", nota, resultado[nota])
        }
    }
}