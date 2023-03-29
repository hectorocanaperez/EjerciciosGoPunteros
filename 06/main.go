package main

import (
	"fmt"
	"strings"
)

func reemplazarVocales(texto *string) {
	vocales := "aeiouAEIOU"
	for _, v := range vocales {
		*texto = strings.ReplaceAll(*texto, string(v), "#")
	}
}

func main() {
	texto := "Estamos haciendo ejercicios de golang"
	fmt.Println("Antes: ", texto)
	reemplazarVocales(&texto)
	fmt.Println("Después: ", texto)
}
