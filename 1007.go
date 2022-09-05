package main

import "fmt"

func main() {
	var a, b, c, d, dif int

	fmt.Printf("Insira o primeiro número:\n")
	fmt.Scanf("%d", &a)
	fmt.Printf("Insira o segundo número:\n")
	fmt.Scanf("%d", &b)
	fmt.Printf("Insira o terceiro número:\n")
	fmt.Scanf("%d", &c)
	fmt.Printf("Insira o terceiro número:\n")
	fmt.Scanf("%d", &d)

	dif = (a * b) - (c * d)

	fmt.Printf("DIFERENCA = %.d\n", dif)
}
