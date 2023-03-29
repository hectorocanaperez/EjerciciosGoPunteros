package main

import "fmt"

func intercambio(a *int, b *int) {
	change := *a
	*a = *b
	*b = change
}
func burbuja(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				intercambio(&slice[j], &slice[j+1])
			}
		}
	}
}

func main() {
	slice := []int{5, 2, 3, 2, 7, 8, 1, 7, 4, 3, 5}
	fmt.Println("Antes de ordenar:", slice)
	burbuja(slice)
	fmt.Println("Después de ordenar:", slice)
}
