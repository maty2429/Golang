package main

import "fmt"

func main() {
	const passwordCorrecto = "secreto123"
	const intentosMaximos = 3

	fmt.Println("Sistema de Login")

	// Usar for como un while
	intentos := 0
	autenticado := false

	for intentos < intentosMaximos && !autenticado {
		var password string
		fmt.Printf("Intento %d/%d - Introduce tu contraseña: ", intentos+1, intentosMaximos)
		// En un programa real usarías fmt.Scanln(&password)

		// Simulamos entradas del usuario
		if intentos == 0 {
			password = "contraseña" // Primer intento incorrecto
		} else if intentos == 1 {
			password = "secret123" // Segundo intento incorrecto
		} else {
			password = "secreto123" // Tercer intento correcto
		}

		fmt.Println(password) // Simular lo que el usuario escribió

		if password == passwordCorrecto {
			autenticado = true
			fmt.Println("✅ Acceso concedido. Bienvenido al sistema.")
			break // Salir del bucle
		} else {
			intentos++
			fmt.Printf("❌ Contraseña incorrecta. Te quedan %d intentos.\n\n", intentosMaximos-intentos)
		}
	}

	if !autenticado {
		fmt.Println("🔒 Cuenta bloqueada. Demasiados intentos fallidos.")
	}
}
