package main

import "fmt"

func main() {
	// Lista de productos con precios
	productos := []struct {
		nombre string
		precio float64
	}{
		{"Laptop", 1200.00},
		{"Teléfono", 800.00},
		{"Auriculares", 150.00},
		{"Monitor", 350.00},
		{"Teclado", 80.00},
	}

	// Calcular precio total
	var total float64
	for i := 0; i < len(productos); i++ {
		total += productos[i].precio
		fmt.Printf("Producto: %s - Precio: $%.2f\n", productos[i].nombre, productos[i].precio)
	}

	fmt.Printf("\nTotal a pagar: $%.2f\n", total)
}
