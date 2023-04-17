package main

import (
	"reflect"
	"testing"
)

func TestBurbuja(t *testing.T) {
	// Prueba 1
	slice1 := []int{5, 2, 3, 2, 7, 8, 1, 7, 4, 3, 5}
	burbuja(slice1)
	expected1 := []int{1, 2, 2, 3, 3, 4, 5, 5, 7, 7, 8}
	if !reflect.DeepEqual(slice1, expected1) {
		t.Errorf("Ordenamiento burbuja incorrecto. Esperado: %v. Obtenido: %v", expected1, slice1)
	}
	// if expected1 == nil {
	// 	t.Errorf("Ordenamiento burbuja incorrecto. Esperado: %v. Obtenido: %v", expected1, slice1)
	// }

	// Prueba 2
	slice2 := []int{9, 3, 7, 1, 8, 4, 2, 6, 5, 10}
	burbuja(slice2)
	expected2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(slice2, expected2) {
		t.Errorf("Ordenamiento burbuja incorrecto. Esperado: %v. Obtenido: %v", expected2, slice2)
	}
}
