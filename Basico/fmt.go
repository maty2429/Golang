package main

import "fmt"

func main() {
	// Imprimir texto simple
	fmt.Println("Hola, Go!") // Imprime con un salto de línea al final
	fmt.Print("Hola, ")      // Imprime sin salto de línea
	fmt.Print("mundo!\n")    // Agrega un salto de línea manualmente

	// Formatear cadenas con Printf
	nombre := "Juan"
	edad := 25
	fmt.Printf("Mi nombre es %s y tengo %d años.\n", nombre, edad) // %s para cadenas, %d para enteros

	// Imprimir valores con formato
	precio := 19.99
	fmt.Printf("El precio es %.2f dólares.\n", precio) // %.2f para mostrar 2 decimales

	// Escanear entrada del usuario
	var entrada string
	fmt.Print("Escribe algo: ")
	fmt.Scanln(&entrada) // Lee una línea de entrada
	fmt.Println("Escribiste:", entrada)

	// Imprimir valores en diferentes formatos
	numero := 255
	fmt.Printf("Decimal: %d, Binario: %b, Hexadecimal: %x\n", numero, numero, numero) // %b para binario, %x para hexadecimal

	// Usar Sprintf para guardar cadenas formateadas
	mensaje := fmt.Sprintf("Hola, %s. Tienes %d mensajes nuevos.", nombre, 5)
	fmt.Println(mensaje)

	// Declarar un slice
	numeros := []int{1, 2, 3, 4, 5}

	// Imprimir el slice directamente
	fmt.Println("Slice:", numeros)

	// Iterar sobre el slice e imprimir cada elemento
	for i, num := range numeros {
		fmt.Printf("Índice %d: %d\n", i, num)
	}

	type Persona struct {
		Nombre string
		Edad   int
	}

	// Crear una instancia del struct
	persona := Persona{
		Nombre: "Juan",
		Edad:   30,
	}

	// Imprimir el struct directamente
	fmt.Println("Persona:", persona)

	// Imprimir los campos del struct
	fmt.Printf("Nombre: %s, Edad: %d\n", persona.Nombre, persona.Edad)
}
