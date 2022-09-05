package main

import "fmt"

func main() {

	var num int
	var horas, valor, salario float64

	fmt.Printf("Insira o número do funcionário:\n")
	fmt.Scanf("%d", &num)
	fmt.Printf("Insira o número de horas trabalhadas desse funcionário:\n")
	fmt.Scanf("%f", &horas)
	fmt.Printf("Insira o valor por horas trabalhadas desse funcionário:\n")
	fmt.Scanf("%f", &valor)

	salario = horas * valor

	fmt.Printf("\n")
	fmt.Printf("NUMBER = %d\n", num)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}
