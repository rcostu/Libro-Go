package main

import "fmt"

func main() {
	var puntero_a_entero *int 	// Declaramos un puntero a entero
	var entero int = 8 			// Y una variable entero que vale 8

	puntero_a_entero = &entero  // puntero\_a\_entero ahora apunta a la variable entero

	fmt.Println(entero)					// 8
	fmt.Println(*puntero_a_entero)		// 8
	fmt.Println(puntero_a_entero)		// 0x420ef000
	fmt.Println(&entero)				// 0x420ef000
}
