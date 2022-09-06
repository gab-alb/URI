package main

import (
	"fmt"
	"math"
)

func main() {

	var x1, y1, x2, y2, distancia float64

	fmt.Printf("Digite os valores de X1 e Y1:2")
	fmt.Scanf("%f%f", &x1, &y1)

	fmt.Printf("Digite os valores de X2 e Y2:")
	fmt.Scanf("%f%f", &x2, y2)

	distancia = math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))

	fmt.Printf("%.4f\n", distancia)

}
