//sin punteros

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	Name string
// 	Apellido string
// 	Edad int
// }

// func main() {
// 	person := Person{
// 		Nombre:    "Hector",
// 		Apellido     "Ocaña",
// 		Edad: 34,
// 	}

// 	jsonData, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println(string(jsonData))
// }

package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func formatoJSON(p *Persona) ([]byte, error) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("error serializando en JSON: %v", err)
	}
	return jsonBytes, nil
}

func main() {
	persona := &Persona{Nombre: "Hector", Apellido: "Ocaña", Edad: 34}
	jsonBytes, err := formatoJSON(persona)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
