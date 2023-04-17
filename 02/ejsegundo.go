package main

import "fmt"

func intercambio(a *int, b *int) {
	change := *a
	*a = *b
	*b = change
}

func main() {
	c := 7
	d := 9
	fmt.Println("Antes del intercambio:", c, d)
	intercambio(&c, &d)
	fmt.Println("Después del intercambio:", c, d)
}
