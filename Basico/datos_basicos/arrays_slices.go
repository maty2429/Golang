package main

import "fmt"

func main() {
	// Declaración de un array de tamaño 5
	var numeros [5]int
	numeros[0] = 10 // Asignar valor al primer elemento
	numeros[1] = 20 // Asignar valor al segundo elemento

	fmt.Println("Array completo:", numeros)     // Imprime el array completo
	fmt.Println("Primer elemento:", numeros[0]) // Accede al primer elemento

	// Declaración e inicialización de un array
	dias := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
	fmt.Println("Días de la semana:", dias)

	// Iterar sobre un array
	for i, dia := range dias {
		fmt.Printf("Día %d: %s\n", i+1, dia)
	}

	// Crear un slice a partir de un array
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4] // Slice desde el índice 1 hasta el 3 (excluye el 4)
	fmt.Println("Slice:", slice)

	// Crear un slice directamente
	numeros2 := []int{10, 20, 30}
	fmt.Println("Slice inicial:", numeros2)

	// Agregar elementos a un slice
	numeros2 = append(numeros2, 40, 50)
	fmt.Println("Slice después de append:", numeros2)

	// Iterar sobre un slice
	for i, num := range numeros2 {
		fmt.Printf("Índice %d: %d\n", i, num)
	}

	// Capacidad y longitud de un slice
	fmt.Println("Longitud del slice:", len(numeros2))
	fmt.Println("Capacidad del slice:", cap(numeros2))
}
