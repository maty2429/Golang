package main

import "fmt"

func main() {

	//El keyword defer se utiliza para posponer la ejecución de una instrucción hasta que la función que la contiene termine.
	fmt.Println("Inicio de la función")

	defer fmt.Println("Esto se ejecuta al final, justo antes de salir de la función") // Se ejecuta al final
	fmt.Println("Ejecutando el cuerpo de la función")

	defer fmt.Println("Este defer se ejecuta en orden inverso al que fue declarado") // Se ejecuta antes del anterior

	//El keyword break se utiliza para salir de un bucle antes de que termine normalmente.
	for i := 1; i <= 10; i++ {
		if i == 5 { // Si el valor de i es 5
			fmt.Println("Se encontró el valor 5, saliendo del bucle")
			break // Sale del bucle
		}
		fmt.Println("Valor:", i)
	}

	//El keyword continue se utiliza para saltar a la siguiente iteración del bucle, omitiendo el resto del código en la iteración actual.
	for i := 1; i <= 10; i++ {
		if i%2 == 0 { // Si el número es par
			continue // Salta a la siguiente iteración
		}
		fmt.Println("Número impar:", i) // Solo imprime números impares
	}
}
