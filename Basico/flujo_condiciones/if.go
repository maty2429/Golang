package main

import "fmt"

func main() {
	numero := -3

	if numero > 0 { // Verifica si el número es mayor que 0
		fmt.Println("El número es positivo")
	} else if numero < 0 { // Verifica si el número es menor que 0
		fmt.Println("El número es negativo")
	} else { // Si no es mayor ni menor, es igual a 0
		fmt.Println("El número es cero")
	}

	numero2 := 15

	if numero2 >= 10 && numero2 <= 20 { // Verifica si el número está entre 10 y 20
		fmt.Println("El número está en el rango de 10 a 20")
	} else {
		fmt.Println("El número está fuera del rango")
	}

	if numero := 8; numero%2 == 0 { // Inicializa la variable y verifica si es par
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}
}
