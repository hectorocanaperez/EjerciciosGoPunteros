package main

import (
	"testing"
)

func TestIntercambio(t *testing.T) {
	// Prueba 1
	a := 5
	b := 7
	intercambio(&a, &b)
	if a != 7 || b != 5 {
		t.Errorf("El intercambio de variables no se realizó correctamente. Esperado: a=7, b=5. Obtenido: a=%d, b=%d", a, b)
	}

	// Prueba 2
	c := 8
	d := 2
	intercambio(&c, &d)
	if c != 2 || d != 8 {
		t.Errorf("El intercambio de variables no se realizó correctamente. Esperado: c=2, d=8. Obtenido: c=%d, d=%d", c, d)
	}
}
