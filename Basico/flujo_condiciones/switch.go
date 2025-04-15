package main

import "fmt"

func main() {
	dia := 3 // Cambia este valor para probar diferentes casos

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6, 7: // Agrupación de casos
		fmt.Println("Fin de semana")
	default: // Caso por defecto
		fmt.Println("Día no válido")
	}
}
