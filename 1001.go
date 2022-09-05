package main

import (
	"fmt"
)

func main() {
	var a, b int
	var x int

	fmt.Printf("Insira o primeiro valor: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("\nInsira o segundo valor: ")
	fmt.Scanf("%d", &b)

	x = a + b

	fmt.Printf("X = %d\n", x)
}
