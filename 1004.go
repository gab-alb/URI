package main

import "fmt"

func main() {
	var a, b, prod int

	fmt.Printf("Insira o primeiro valor: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Insira o primeiro valor: ")
	fmt.Scanf("%d", &b)

	prod = (a * b)

	fmt.Printf("PROD = %d\n", prod)

}
