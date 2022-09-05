package main

import "fmt"

func main() {
	var n1, n2, media float64

	fmt.Printf("Insira a primeira nota do aluno:\n")
	fmt.Scanf("%f", &n1)
	fmt.Printf("Insira a segunda nota do aluno:\n")
	fmt.Scanf("%f", &n2)

	media = ((n1 * 3.5) + (n2 * 7.5)) / 11

	fmt.Printf("MEDIA = %.5f\n", media)

}
