package main

import "fmt"

func main() {
	var km, valor, consumo float64

	fmt.Printf("Indique a kilometragem rodada pelo veiculo:")
	fmt.Scanf("%f", &km)
	fmt.Printf("Indique o consumo em Litros:")
	fmt.Scanf("%f", &consumo)

	valor = km / consumo

	fmt.Printf("%.3f km/l\n", valor)

}
