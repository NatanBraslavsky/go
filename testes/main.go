package main

import "fmt"

func main() {
    var valor int

    notas := []int{100, 50, 20, 10, 5, 2, 1}
    resultado := make(map[int]int)

    fmt.Print("Digite um valor: ")
    fmt.Scan(&valor)

    for _, nota := range notas {
        if valor >= nota {
            quantidade := valor / nota
            resultado[nota] = quantidade
            valor = valor % nota
        }
    }

    for _, nota := range notas {
        if resultado[nota] > 0 {
            fmt.Printf("%d nota(s) de R$ %d\n", resultado[nota], nota)
        }
    }
}
