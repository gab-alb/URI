package main

import "fmt"

func main() {
	var a, b, soma int

	fmt.Printf("Insira o primeiro valor: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Insira o primeiro valor: ")
	fmt.Scanf("%d", &b)

	soma = (a + b)

	fmt.Printf("SOMA = %d", soma)
}
