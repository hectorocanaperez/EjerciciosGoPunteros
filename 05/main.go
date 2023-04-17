package main

import (
	"fmt"
	"sort"
)

func ArrayOrdenado(array *[]int) []*int {
	numero := len(*array)
	ptrs := make([]*int, numero)
	for i := 0; i < numero; i++ {
		ptrs[i] = &(*array)[i]
	}
	sort.Slice(ptrs, func(i, j int) bool {
		return *ptrs[i] < *ptrs[j]
	})
	return ptrs
}

func main() {
	array := []int{3, 11, 14, 21, 35, 9, 12, 6, 35, 3, 5}
	fmt.Println("Antes:", array)
	ptrs := ArrayOrdenado(&array)
	fmt.Println("Despues:")
	for i := 0; i < len(ptrs); i++ {
		fmt.Printf("%d ", *ptrs[i])
	}
}
