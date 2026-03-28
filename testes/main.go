package main

import "fmt"

import "strings"

func verificarPalindromo(txt string) bool {
    arr := strings.Split(txt, "") // separa por caractere

    inv := []string{}

    for i := len(arr) - 1; i >= 0; i-- {
        inv = append(inv, arr[i])
    }

    for i := 0; i < len(arr); i++ { // corrigido aqui
        if arr[i] != inv[i] {
            return false
        }
    }

    return true
}

func main() {
	fmt.Println(verificarPalindromo("atan"))
}

