package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Printf("Insira o valor de A.")
	fmt.Scanf("%d", &a)
	fmt.Printf("Insira o valor de B.")
	fmt.Scanf("%d", &b)
	fmt.Printf("Insira o valor de B.")
	fmt.Scanf("%d", &c)

	if a > b && a > c {
		fmt.Printf("%d eh maior\n", a)
	} else if b > c {
		fmt.Printf("%d eh maior\n", b)
	} else {
		fmt.Printf("%d eh maior\n", c)
	}

}
