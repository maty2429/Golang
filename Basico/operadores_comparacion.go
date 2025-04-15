package main

import "fmt"

func main() {
	//operador de comparacion

	fmt.Println(4 > 2)  // 4 es mayor que 2
	fmt.Println(4 < 2)  // 4 es menor que 2
	fmt.Println(4 >= 2) // 4 es mayor o igual que 2
	fmt.Println(4 <= 2) // 4 es menor o igual que 2
	fmt.Println(4 == 2) // 4 es igual que 2
	fmt.Println(4 != 2) // 4 es diferente que 2

	//operador logico
	fmt.Println(true && true)   // true , debe cumplirse las dos condiciones si uno es falso el resultado es falso
	fmt.Println(true && false)  // false
	fmt.Println(true || true)   // true , debe cumplirse una de las dos condiciones si las dos son falsas el resultado es falso
	fmt.Println(true || false)  // true, debe cumplirse una de las dos condiciones si las dos son falsas el resultado es falso pero si una es verdadera el resultado es verdadero
	fmt.Println(false || false) // false

	//operador de negacion
	fmt.Println(!true) // false, estoy negando el valor de true

}
