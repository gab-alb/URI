package main

import "fmt"

func main() {

	var nome string
	var salario, vendas, total float64

	fmt.Printf("Insira o nome do funcionário:\n")
	fmt.Scanf("%s", &nome)
	fmt.Printf("Insira o salário desse funcionário:\n")
	fmt.Scanf("%f", &salario)
	fmt.Printf("Insira o valor das vendas desse funcionário:\n")
	fmt.Scanf("%f", &vendas)

	total = salario + (vendas * 0.15)

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
