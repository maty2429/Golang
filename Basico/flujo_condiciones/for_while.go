package main

import "fmt"

func main() {
	// La estructura completa del bucle for en Go tiene tres componentes principales:
	// for [inicialización]; [condición]; [post-instrucción] { ... }

	fmt.Println("Explicación detallada del bucle for:")

	/*
	 * ESTRUCTURA DEL FOR COMPLETO:
	 * ----------------------------
	 * for i := 1; i <= 10; i++ { ... }
	 *     |       |       |
	 *     |       |       └── Post-instrucción: Se ejecuta después de cada iteración
	 *     |       |
	 *     |       └── Condición: Determina si el bucle continúa (true) o termina (false)
	 *     |
	 *     └── Inicialización: Se ejecuta una sola vez al inicio
	 */

	fmt.Println("\n1. Simulación de carga de archivos:")
	for i := 1; i <= 5; i++ {
		// i := 1 -> Inicialización: Declara e inicializa el contador en 1
		// i <= 5 -> Condición: Continúa el bucle mientras i sea menor o igual a 5
		// i++ -> Post-instrucción: Incrementa i en 1 después de cada iteración

		// Simulamos que cada archivo tiene un tamaño diferente
		tamanoArchivo := i * 20 // Tamaño en MB

		// Simulamos un progreso de carga para cada archivo
		fmt.Printf("Cargando archivo %d (%dMB): ", i, tamanoArchivo)

		// Bucle anidado para mostrar el progreso
		for progreso := 0; progreso <= 100; progreso += 25 {
			// progreso := 0 -> Inicializa el porcentaje de progreso en 0%
			// progreso <= 100 -> Continúa hasta llegar al 100%
			// progreso += 25 -> Incrementa en saltos de 25% cada vez

			fmt.Printf("%d%% ", progreso)
		}
		fmt.Println("¡Completado!")
	}

	fmt.Println("\n2. Fibonacci con límite de números:")
	// Otro ejemplo: Secuencia de Fibonacci con límite de valores
	/*
	 * Este ejemplo muestra que puedes inicializar múltiples variables,
	 * tener una condición compleja, y actualizar múltiples valores
	 * en la post-instrucción.
	 */
	for a, b := 0, 1; a < 100; a, b = b, a+b {
		// a, b := 0, 1 -> Inicialización múltiple: a empieza en 0, b en 1
		// a < 100 -> Condición: Continúa mientras a sea menor que 100
		// a, b = b, a+b -> Post-instrucción: Actualiza ambos valores simultáneamente
		//                  (a toma el valor de b, y b toma el valor de a+b)

		fmt.Printf("%d ", a)
	}
	fmt.Println()

	fmt.Println("\n3. Contador con diferentes incrementos:")
	/*
	 * Este ejemplo muestra cómo puedes personalizar los incrementos
	 * y usar lógica más compleja en la sección de post-instrucción
	 */
	for i := 1; i <= 20; i = i * 2 {
		// i := 1 -> Inicialización: Comienza con 1
		// i <= 20 -> Condición: Continúa mientras i sea menor o igual a 20
		// i = i*2 -> Post-instrucción: Multiplica i por 2 en cada iteración

		fmt.Printf("%d ", i) // Imprime: 1, 2, 4, 8, 16
	}
	fmt.Println()

	fmt.Println("\n4. Bucle con condición compleja:")
	/*
	 * En este ejemplo, la condición combina múltiples variables
	 * para determinar si el bucle debe continuar
	 */
	for intentos, totalPuntos := 1, 0; intentos <= 3 && totalPuntos < 100; intentos++ {
		// Inicialización doble: intentos = 1, totalPuntos = 0
		// Condición compuesta: Continúa si AMBAS condiciones son verdaderas
		// Post-instrucción: Solo incrementa intentos

		// Simulamos puntos ganados en cada intento
		puntosGanados := intentos * 25
		totalPuntos += puntosGanados

		fmt.Printf("Intento %d: Ganaste %d puntos. Total: %d\n",
			intentos, puntosGanados, totalPuntos)

		// Nota: Este bucle termina después de 3 intentos O si alcanzas 100 puntos
	}

	fmt.Println("\n5. Declaración de variables fuera del bucle:")
	/*
	 * También puedes omitir la inicialización si ya tienes variables
	 * declaradas e inicializadas previamente
	 */
	contador := 5
	for ; contador > 0; contador-- {
		// La inicialización se omite (nota el ; inicial)
		// contador > 0 -> Continúa mientras contador sea mayor que 0
		// contador-- -> Decrementa contador en cada iteración

		fmt.Printf("Cuenta regresiva: %d\n", contador)
	}

	fmt.Println("\n6. Omitiendo la post-instrucción:")
	/*
	 * Puedes omitir la post-instrucción y manejarla dentro del cuerpo del bucle
	 */
	x := 1
	for x <= 16 {
		// x := 1 -> Variable ya inicializada
		// x <= 16 -> Condición: Continúa mientras x sea menor o igual a 16
		// Post-instrucción omitida (nota el ; final)

		fmt.Printf("%d ", x)
		x = x * 2 // La actualización se hace dentro del bucle
	}
	fmt.Println()

	fmt.Println("\n7. Bucle for como while:")
	/*
	 * Si solo necesitas la condición, puedes usar for como un while
	 * omitiendo tanto la inicialización como la post-instrucción
	 */
	suma := 0
	num := 1
	for suma < 50 {
		// Equivalente a "while (suma < 50)" en otros lenguajes
		suma += num
		fmt.Printf("Número: %d, Suma acumulada: %d\n", num, suma)
		num++
	}

	fmt.Println("\n8. Bucle infinito con break:")
	/*
	 * También puedes crear un bucle infinito y utilizar
	 * break para salir cuando se cumpla alguna condición
	 */
	contador = 1
	for {
		// Bucle infinito (no hay condición)
		fmt.Printf("Iteración %d\n", contador)

		if contador >= 5 {
			fmt.Println("¡Condición de salida alcanzada!")
			break // Sale del bucle
		}

		contador++
	}
}
