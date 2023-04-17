package main

import (
	"testing"
)

func TestEncontrarMayor(t *testing.T) {
	slice := []int{33, 51, 4, 11, 15, 12, 22, 6, 40, 3}
	expected := 51
	result := *encontrarMayor(slice)

	if result != expected {
		t.Errorf("El resultado fue %d pero se esperaba %d", result, expected)
	}
}
