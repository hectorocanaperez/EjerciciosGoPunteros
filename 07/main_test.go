package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFormatoJSON(t *testing.T) {
	persona := &Persona{Nombre: "Hector", Apellido: "Ocaña", Edad: 34}
	expectedJSON := []byte(`{"Nombre":"Hector","Apellido":"Ocaña","Edad":34}`)

	jsonBytes, err := formatoJSON(persona)
	if err != nil {
		t.Errorf("Se esperaba un JSON válido pero se produjo un error: %v", err)
	}

	if !reflect.DeepEqual(jsonBytes, expectedJSON) {
		t.Errorf("El resultado fue %s pero se esperaba %s", jsonBytes, expectedJSON)
	}

	// Validamos que el JSON sea válido y se pueda deserializar en una estructura Persona
	var p Persona
	err = json.Unmarshal(jsonBytes, &p)
	if err != nil {
		t.Errorf("El JSON generado no es válido: %v", err)
	}
}
