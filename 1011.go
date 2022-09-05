package main

import "fmt"

func main() {
	var volume, raio float64

	fmt.Printf("Qual o raio da esfera:")
	fmt.Scanf("%f", &raio)

	pi := 3.14159

	volume = pi * (raio * raio * raio) * (4 / 3)

	fmt.Printf("VOLUME = %.3f\n", volume)

}
