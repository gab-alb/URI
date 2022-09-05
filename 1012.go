package main

import "fmt"

func main() {
	var a, b, c, tri, cir, tra, qua, ret, pi float64

	fmt.Printf("Insira o valor de A.")
	fmt.Scanf("%f", &a)
	fmt.Printf("Insira o valor de B.")
	fmt.Scanf("%f", &b)
	fmt.Printf("Insira o valor de C.")
	fmt.Scanf("%f", &c)

	pi = 3.14159

	tri = (a * c) / 2
	cir = pi * (c * c)
	tra = ((a + b) * c) / 2
	qua = b * b
	ret = a * b

	fmt.Printf("TRIANGULO: %.3f\n", tri)
	fmt.Printf("CIRCULO: %.3f\n", cir)
	fmt.Printf("TRAPEZIO: %.3f\n", tra)
	fmt.Printf("QUADRADO: %.3f\n", qua)
	fmt.Printf("RETANGULO: %.3f\n", ret)
}
