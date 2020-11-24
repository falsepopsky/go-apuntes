package main

import (
	"fmt"
)

// package level, las variables que no se usan en este scope se permiten correr
var (
	// luego del nombre de la variable me permite definir el tipo
	nombreActor  string = "Elisabeth Sladen"
	compania     string = "Sarah"
	numeroDoctor int    = 3
	temporada    int    = 11
)

var (
	counter int = 0
)

func main() {
	fmt.Println(numeroDoctor)
	// puedo pisar la variable fuera del scope, excepto con :=
	var numeroDoctor = 5
	fmt.Printf("%v, %T", numeroDoctor, numeroDoctor)

	fmt.Printf("%v, %T", nombreActor, nombreActor)

	// cambio el tipo
	var nuevoTipoNumeroDoctor float32
	nuevoTipoNumeroDoctor = float32(numeroDoctor)
	fmt.Printf("%v, %T", nuevoTipoNumeroDoctor, nuevoTipoNumeroDoctor)
}
