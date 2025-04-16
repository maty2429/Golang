package main

import "fmt"

func main() {
	archivos := []string{"documento.txt", "imagen.jpg", "video.mp4", "malware.exe", "datos.csv"}

	fmt.Println("Escaneando archivos...")

	// Usando for con range para iterar sobre slice
	for i, archivo := range archivos {
		fmt.Printf("Procesando archivo %d: %s\n", i+1, archivo)

		// Simular detección de virus
		if archivo == "malware.exe" {
			fmt.Println("⚠️ ¡ALERTA! Virus detectado en:", archivo)
			continue // Saltar este archivo y seguir con el siguiente
		}

		// Simular procesamiento
		extension := archivo[len(archivo)-3:]
		switch extension {
		case "txt", "csv":
			fmt.Println("  - Archivo de texto procesado")
		case "jpg", "png":
			fmt.Println("  - Imagen procesada")
		case "mp4", "mov":
			fmt.Println("  - Video procesado")
		default:
			fmt.Println("  - Archivo de tipo desconocido")
		}
	}

	fmt.Println("Procesamiento finalizado")
}
