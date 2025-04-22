package main

import (
	"fmt"
	"time"
)

func main() {
	// Datos de transacciones simuladas
	transacciones := []struct {
		fecha    string
		tipo     string
		monto    float64
		aprobada bool
	}{
		{"2025-04-10", "depósito", 1500.00, true},
		{"2025-04-11", "retiro", 200.00, true},
		{"2025-04-12", "retiro", 5000.00, false},
		{"2025-04-13", "depósito", 350.00, true},
		{"2025-04-14", "transferencia", 400.00, true},
	}

	saldo := 1000.00 // Saldo inicial

	fmt.Printf("Estado de cuenta - Fecha actual: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("Saldo inicial: $%.2f\n\n", saldo)
	fmt.Println("Historial de transacciones:")

	// Procesar transacciones
	for i, t := range transacciones {
		fmt.Printf("%d. [%s] %s por $%.2f - ", i+1, t.fecha, t.tipo, t.monto)

		if !t.aprobada {
			fmt.Println("RECHAZADA")
			continue
		}

		// Actualizar saldo según el tipo de transacción
		switch t.tipo {
		case "depósito":
			saldo += t.monto
			fmt.Println("COMPLETADA")
		case "retiro", "transferencia":
			if saldo >= t.monto {
				saldo -= t.monto
				fmt.Println("COMPLETADA")
			} else {
				fmt.Println("RECHAZADA (fondos insuficientes)")
			}
		}

		fmt.Printf("   Nuevo saldo: $%.2f\n", saldo)
	}

	fmt.Printf("\nSaldo actual: $%.2f\n", saldo)
}
