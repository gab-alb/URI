package main

import "fmt"

func main() {

	var n1, n2, n3, media float64

	fmt.Printf("Insira a primeira nota do aluno:\n")
	fmt.Scanf("%f", &n1)
	fmt.Printf("Insira a segunda nota do aluno:\n")
	fmt.Scanf("%f", &n2)
	fmt.Printf("Insira a terceira nota do aluno:\n")
	fmt.Scanf("%f", &n3)

	media = ((n1 * 2) + (n2 * 3) + (n3 * 5)) / 10

	fmt.Printf("MEDIA = %.1f\n", media)

}
