package main

import (
	"fmt"
	"math"
)

func main() {

	var raio float64
	pi := 3.14159

	fmt.Println("Insira o valor do Raio:")
	fmt.Scanf("%f", &raio)

	area := (pi * (math.Pow(raio, 2)))

	fmt.Printf("A=%.4f\n", area)

}
