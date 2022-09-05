package main

import "fmt"

func main() {
	var cod1, cod2 int
	var valor1, valor2, total, num1, num2 float64

	fmt.Printf("Insira o código do primeiro produto, a quantidade desse produto e seu valor:")
	fmt.Scanf("%d%d%f", &cod1, &num1, &valor1)
	fmt.Printf("Insira o código do segundo produto, a quantidade desse produto e seu valor:")
	fmt.Scanf("%d%d%f", &cod2, &num2, &valor2)

	total = (num1 * valor1) + (num2 * valor2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
