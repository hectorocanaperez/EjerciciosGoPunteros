package main

import "testing"

func TestReemplazarVocales(t *testing.T) {
	texto := "Estamos haciendo ejercicios de golang"
	expected := "##st#m#s h#c##ndo #j#rc#c#s d# g#l#ng"
	reemplazarVocales(&texto)

	if texto == expected {
		t.Errorf("El resultado fue %s pero se esperaba %s", texto, expected)
	}
}
