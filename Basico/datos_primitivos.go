package main

import "fmt"

func main() {
	// Tipos numéricos enteros
	var entero8 int8 = 127                   // Entero de 8 bits
	var entero16 int16 = 32767               // Entero de 16 bits
	var entero32 int32 = 2147483647          // Entero de 32 bits
	var entero64 int64 = 9223372036854775807 // Entero de 64 bits
	var entero int = 42                      // Entero (tamaño depende de la arquitectura)

	// Tipos numéricos sin signo
	var uEntero8 uint8 = 255                    // Entero sin signo de 8 bits
	var uEntero16 uint16 = 65535                // Entero sin signo de 16 bits
	var uEntero32 uint32 = 4294967295           // Entero sin signo de 32 bits
	var uEntero64 uint64 = 18446744073709551615 // Entero sin signo de 64 bits
	var uEntero uint = 42                       // Entero sin signo (tamaño depende de la arquitectura)

	// Tipos de punto flotante
	var flotante32 float32 = 3.14              // Flotante de 32 bits
	var flotante64 float64 = 3.141592653589793 // Flotante de 64 bits

	// Tipo complejo (números complejos)
	var complejo64 complex64 = 1 + 2i   // Número complejo de 64 bits
	var complejo128 complex128 = 1 + 2i // Número complejo de 128 bits

	// Tipo booleano
	var booleano bool = true // Valor booleano (true o false)

	// Tipo string
	var cadena string = "Hola, Go!" // Cadena de texto

	// Tipo byte (alias de uint8)
	var caracter byte = 'A' // Representa un carácter en ASCII

	// Tipo rune (alias de int32, para caracteres Unicode)
	var simbolo rune = '✓' // Representa un carácter Unicode

	// Imprimir los valores
	fmt.Println("Entero8:", entero8)
	fmt.Println("Entero16:", entero16)
	fmt.Println("Entero32:", entero32)
	fmt.Println("Entero64:", entero64)
	fmt.Println("Entero (arquitectura):", entero)
	fmt.Println("uEntero8:", uEntero8)
	fmt.Println("uEntero16:", uEntero16)
	fmt.Println("uEntero32:", uEntero32)
	fmt.Println("uEntero64:", uEntero64)
	fmt.Println("uEntero (arquitectura):", uEntero)
	fmt.Println("Flotante32:", flotante32)
	fmt.Println("Flotante64:", flotante64)
	fmt.Println("Complejo64:", complejo64)
	fmt.Println("Complejo128:", complejo128)
	fmt.Println("Booleano:", booleano)
	fmt.Println("Cadena:", cadena)
	fmt.Println("Byte (carácter):", caracter)
	fmt.Println("Rune (símbolo):", simbolo)
}
