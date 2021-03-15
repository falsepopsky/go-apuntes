package main

import "fmt"

/*
Package Level
=============

En este nivel se puden declarar variables que puedo utilizar o no y al momento de compilación, no me arroja error.
Las variables que se declaran y no se utilicen en otro scope que no sea el package, al momento de compilar arroja error.
*/

var (
	juego string = "Silent Hill"
	compania     string = "Konami"
	fechaDeLanzamiento int = 1999
	creador string
)


func main() {
	fmt.Printf("%T\n", juego) // %T muestra el tipo de dato.
	fmt.Printf("%v\n", compania) // %v muestra el valor.

	// Iniciando la variable que fue declarado a nivel del package

	creador = "Keiichiro Toyama"
	fmt.Println(creador)

	/*  
	las variables no se pueden redeclarar, dentro de un mismo scope sin embargo,
	cuando es una variable a nivel de package se permite (no funciona con la asignación corta := ).
	*/

	fmt.Println(fechaDeLanzamiento)
	fechaDeLanzamiento = 2021
	fmt.Println(fechaDeLanzamiento)
	fechaDeLanzamiento = 1900
	fmt.Println(fechaDeLanzamiento)
	
	otroJuego := "Metal Gear Solid"
	fmt.Println(otroJuego)

	piNumber := 3.14159265359
	fmt.Printf("%T\n", piNumber)
}
