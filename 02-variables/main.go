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
	compania string = "Konami"
	fechaDeLanzamiento int = 1999
	creador string
)


func main() {
	fmt.Printf("El tipo de dato de la variable juego es: %T\n", juego) // %T muestra el tipo de dato.
	fmt.Printf("El valor de la variable compania es: %v\n", compania) // %v muestra el valor.

	creador = "Keiichiro Toyama" // Iniciando la variable que fue declarado a nivel del package
	fmt.Println(creador)

	/*  
	las variables no se pueden redeclarar dentro de un mismo scope sin embargo,
	cuando es una variable a nivel de package se permite (no funciona con la asignación corta := ).
	*/

	fmt.Println(fechaDeLanzamiento) // fecha original
	fechaDeLanzamiento = 2021
	fmt.Println(fechaDeLanzamiento) // fecha modificada

	var texto string = juego + " fue creado por " + creador + " y publicado por la compañia japonesa: " + compania
	fmt.Println(texto)

	otroJuego := "Metal Gear Solid"
	fmt.Println(otroJuego)

	piNumber := 3.14159265359
	fmt.Printf("%T\n", piNumber)
}
