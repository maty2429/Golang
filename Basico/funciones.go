package main

import (
	"fmt"
	"strings"
)

// Función simple sin parámetros ni retorno
func saludar() {
	fmt.Println("¡Hola, mundo!")
}

// Función con parámetros
func sumar(a int, b int) int {
	return a + b
}

// Función con múltiples valores de retorno
func dividir(dividendo, divisor int) (int, int) {
	cociente := dividendo / divisor
	resto := dividendo % divisor
	return cociente, resto
}

// Función con parámetros nombrados y valores por defecto (simulado)
func saludarConNombre(nombre string) string {
	if nombre == "" {
		nombre = "Invitado"
	}
	return fmt.Sprintf("¡Hola, %s!", nombre)
}

func greet(firstname, lastname string) { //esto es para que reciba dos parametros de tipo string
	fmt.Println("Hola", firstname, lastname)
}

func convert(text string) (string, string) { //le digo que retorna dos strings
	lower := strings.ToLower(text)
	upper := strings.ToUpper(text)

	return lower, upper
}

func filter(nums []int, callback func(int) bool) []int { //va a recibir 2 pqrametros, un slice de enteros y una funcion que recibe un entero y retorna un booleano, y como retorno va a devolver un slice de enteros
	result := make([]int, 0, len(nums)) //esto es para crear un slice de enteros vacio

	for _, v := range nums {
		//esto es para recorrer el slice de enteros
		if callback(v) { //esto es para que si la funcion callback retorna true, entonces se va a agregar el valor al slice de enteros
			result = append(result, v)
		}
	}
	return result
}

func sum(a int) func(int) int { //esto es para que me retorne una funcion que recibe un entero y retorna un entero
	return func(b int) int { //esto es para que me retorne un entero
		return a + b
	}
}

// Función variádica (acepta un número variable de argumentos)
func sumarTodos(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// Función anónima (usada como variable)
var multiplicar = func(a, b int) int {
	return a * b
}

// Función que utiliza otra función como parámetro
func operar(a, b int, operacion func(int, int) int) int {
	return operacion(a, b)
}

// Función recursiva
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Llamar a una función simple
	saludar()

	// Llamar a una función con parámetros
	fmt.Println("Suma:", sumar(5, 3))

	// Llamar a una función con múltiples valores de retorno
	cociente, resto := dividir(10, 3)
	fmt.Printf("División: Cociente = %d, Resto = %d\n", cociente, resto)

	// Llamar a una función con parámetros nombrados
	fmt.Println(saludarConNombre("Juan"))
	fmt.Println(saludarConNombre(""))

	// Llamar a una función variádica
	fmt.Println("Suma de todos:", sumarTodos(1, 2, 3, 4, 5))

	// Usar una función anónima
	fmt.Println("Multiplicación:", multiplicar(4, 5))

	// Usar una función como parámetro
	resultado := operar(6, 2, func(a, b int) int {
		return a - b
	})
	fmt.Println("Operación personalizada (resta):", resultado)

	// Llamar a una función recursiva
	fmt.Println("Factorial de 5:", factorial(5))
}
