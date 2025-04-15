package main

import "fmt"

func main() {
	// Declaración explícita con tipo
	var entero int = 10             // Variable de tipo entero
	var flotante float64 = 3.14     // Variable de tipo flotante
	var cadena string = "Hola, Go!" // Variable de tipo cadena
	var booleano bool = true        // Variable de tipo booleano

	// Declaración implícita (Go infiere el tipo)
	numero := 42              // Go infiere que es un int
	mensaje := "Hola, mundo!" // Go infiere que es un string
	esVerdadero := false      // Go infiere que es un bool

	// Declaración de múltiples variables
	var x, y, z int = 1, 2, 3     // Declaración de varias variables del mismo tipo
	a, b, c := 4.5, "texto", true // Declaración implícita de múltiples variables con diferentes tipos

	// Variables sin inicializar (toman el valor cero de su tipo)
	var sinValorInt int       // Valor por defecto: 0
	var sinValorFloat float64 // Valor por defecto: 0.0
	var sinValorString string // Valor por defecto: ""
	var sinValorBool bool     // Valor por defecto: false

	// Constantes (no son variables, pero son importantes)
	const constante = 100      // Constante de tipo int
	const pi float64 = 3.14159 // Constante de tipo flotante

	// Declaración de constantes en bloque
	const (
		limite   = 500
		mensaje2 = "Constante en bloque"
	)

	// Declaración de variables en bloque
	var (
		edad    int    = 25
		nombre  string = "Juan"
		activo  bool   = true
		salario float64
	)

	// Constantes calculadas
	const (
		base   = 10
		altura = 5
		area   = base * altura // Constante calculada
	)

	// Imprimir las variables y constantes adicionales
	fmt.Println("Límite:", limite)
	fmt.Println("Mensaje 2:", mensaje2)
	fmt.Println("Edad:", edad)
	fmt.Println("Nombre:", nombre)
	fmt.Println("Activo:", activo)
	fmt.Println("Salario (sin inicializar):", salario)
	fmt.Println("Área calculada:", area)

	// Imprimir las variables originales
	fmt.Println("Entero:", entero)
	fmt.Println("Flotante:", flotante)
	fmt.Println("Cadena:", cadena)
	fmt.Println("Booleano:", booleano)
	fmt.Println("Número:", numero)
	fmt.Println("Mensaje:", mensaje)
	fmt.Println("Es verdadero:", esVerdadero)
	fmt.Println("Múltiples variables:", x, y, z)
	fmt.Println("Múltiples tipos:", a, b, c)
	fmt.Println("Sin valor (int):", sinValorInt)
	fmt.Println("Sin valor (float):", sinValorFloat)
	fmt.Println("Sin valor (string):", sinValorString)
	fmt.Println("Sin valor (bool):", sinValorBool)
	fmt.Println("Constante:", constante)
	fmt.Println("Pi:", pi)
}
