package main

import "fmt"

func main() {
	// Declaración de variables
	var a int = 10
	var b int = 3

	// Suma
	suma := a + b // Suma de a y b
	fmt.Println("Suma:", suma)

	// Resta
	resta := a - b // Resta de a y b
	fmt.Println("Resta:", resta)

	// Multiplicación
	multiplicacion := a * b // Multiplicación de a y b
	fmt.Println("Multiplicación:", multiplicacion)

	// División
	division := a / b // División entera de a entre b
	fmt.Println("División (entera):", division)

	// Módulo (resto de la división)
	modulo := a % b // Resto de la división de a entre b
	fmt.Println("Módulo:", modulo)

	// Operaciones con flotantes
	var x float64 = 10.0
	var y float64 = 3.0

	// División con números flotantes
	divisionFlotante := x / y // División de x entre y
	fmt.Println("División (flotante):", divisionFlotante)

	// Operaciones combinadas
	resultado := (a + b) * (a - b) // Ejemplo de operación combinada
	fmt.Println("Resultado de operación combinada:", resultado)
}
