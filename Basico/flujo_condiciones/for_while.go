package main

import "fmt"

func main() {
	// Ciclo for con condición
	num := 1
	for num <= 5 { // Mientras num sea menor o igual a 5
		fmt.Println("Número:", num)
		num++ // Incrementa num en 1
	}

	// Ciclo for con break y continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // Si el número es par
			continue // Salta a la siguiente iteración
		}
		if i > 7 { // Si el número es mayor que 7
			break // Termina el ciclo
		}
		fmt.Println("Número impar:", i)
	}

	// Ciclo for anidado
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j) // Imprime la multiplicación
		}
	}
}
