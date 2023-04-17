package main

import (
	"testing"
)

func TestSuma(t *testing.T) {
	var resultado int
	suma(3, 6, &resultado)
	if resultado != 9 {
		t.Errorf("La suma esperada era 9, pero se obtuvo %d", resultado)
	}
}

func TestSuma1(t *testing.T) {
	resultado := suma1(3, 6)
	if *resultado != 9 {
		t.Errorf("La suma esperada era 9, pero se obtuvo %d", *resultado)
	}
}
