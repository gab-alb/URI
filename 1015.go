package main

import (
	"fmt"
	"math"
)

func main() {

	var x1, y1, x2, y2, distancia float64

	fmt.Printf("Digite os valores de X1:")
	fmt.Scanf("%f", &x1)
	fmt.Printf("Digite os valores de Y1:")
	fmt.Scanf("%f", &y1)
	fmt.Printf("Digite os valores de X2:")
	fmt.Scanf("%f", &x2)
	fmt.Printf("Digite os valores de Y2:")
	fmt.Scanf("%f", &y2)

	distancia = math.Sqrt(((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))

	fmt.Printf("%.4f\n", distancia)

}
