package main

import "fmt"

func main() {
	// Operador AND lógico (&&)
	fmt.Println("Operador AND lógico:")
	fmt.Println(true && true)   // true: ambas condiciones son verdaderas
	fmt.Println(true && false)  // false: una condición es falsa
	fmt.Println(false && false) // false: ambas condiciones son falsas

	// Operador OR lógico (||)
	fmt.Println("\nOperador OR lógico:")
	fmt.Println(true || true)   // true: al menos una condición es verdadera
	fmt.Println(true || false)  // true: al menos una condición es verdadera
	fmt.Println(false || false) // false: ambas condiciones son falsas

	// Operador NOT lógico (!)
	fmt.Println("\nOperador NOT lógico:")
	fmt.Println(!true)  // false: invierte el valor de true
	fmt.Println(!false) // true: invierte el valor de false

	// Ejemplo combinado de operadores lógicos
	fmt.Println("\nEjemplo combinado:")
	edad := 25
	ingreso := 3000
	fmt.Println(edad > 18 && ingreso > 2000) // true: ambas condiciones son verdaderas
	fmt.Println(edad > 18 || ingreso > 5000) // true: al menos una condición es verdadera
	fmt.Println(!(edad > 18))                // false: invierte el resultado de la condición
}
