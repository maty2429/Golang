package main

import (
	"fmt"
	"math"
)

func main() {
	// Datos de temperatura diaria de una semana
	temperaturas := []float64{22.5, 25.3, 27.8, 24.1, 20.5, 21.2, 23.7}

	// Cálculo de estadísticas básicas
	var suma float64
	var max = temperaturas[0]
	var min = temperaturas[0]

	fmt.Println("Análisis de temperaturas diarias:")

	for i, temp := range temperaturas {
		dia := i + 1
		fmt.Printf("Día %d: %.1f°C\n", dia, temp)

		suma += temp

		if temp > max {
			max = temp
		}
		if temp < min {
			min = temp
		}
	}

	promedio := suma / float64(len(temperaturas))

	// Cálculo de desviación estándar
	var sumaDiferenciasCuadrado float64
	for _, temp := range temperaturas {
		diferencia := temp - promedio
		sumaDiferenciasCuadrado += diferencia * diferencia
	}
	desviacionEstandar := math.Sqrt(sumaDiferenciasCuadrado / float64(len(temperaturas)))

	fmt.Printf("\nEstadísticas:\n")
	fmt.Printf("- Temperatura promedio: %.2f°C\n", promedio)
	fmt.Printf("- Temperatura máxima: %.2f°C\n", max)
	fmt.Printf("- Temperatura mínima: %.2f°C\n", min)
	fmt.Printf("- Desviación estándar: %.2f°C\n", desviacionEstandar)
}
