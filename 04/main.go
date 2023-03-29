package main

import "fmt"

func encontrarMayor(slice []int) *int {

	max := &slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > *max {
			max = &slice[i]
		}
	}
	return max
}

func main() {
	slice := []int{33, 51, 4, 11, 15, 12, 22, 6, 40, 3}
	max := encontrarMayor(slice)

	fmt.Println("El número más grande es", *max)
	fmt.Println("Su dirección de memoria es", &max)

}
