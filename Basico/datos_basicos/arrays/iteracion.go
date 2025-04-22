package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("===== EJEMPLOS DETALLADOS DE ITERACIÓN DE ARRAYS EN GO =====")

	// ==========================================
	// EJEMPLO 1: ITERACIÓN BÁSICA DE UN ARRAY UNIDIMENSIONAL
	// ==========================================
	fmt.Println("\n1. ITERACIÓN BÁSICA DE UN ARRAY UNIDIMENSIONAL")

	// Declaración e inicialización de un array de frutas
	frutas := [5]string{"Manzana", "Naranja", "Plátano", "Pera", "Uva"}
	fmt.Printf("Array original: %v\n", frutas)

	// Método 1: Bucle for tradicional con contador
	fmt.Println("\n1.1. Iteración con bucle for tradicional:")
	for i := 0; i < len(frutas); i++ {
		// i es el índice, va desde 0 hasta len(frutas)-1
		fmt.Printf("  frutas[%d] = %s\n", i, frutas[i])

		// Podemos modificar el array durante la iteración
		if frutas[i] == "Plátano" {
			frutas[i] = "Banana" // Cambiamos Plátano por Banana
		}
	}
	fmt.Printf("  Array después de modificar: %v\n", frutas)

	// Método 2: Iteración con range (devuelve índice y valor)
	fmt.Println("\n1.2. Iteración con range (índice y valor):")
	for i, fruta := range frutas {
		// i = índice, fruta = valor (una copia del valor en el array)
		fmt.Printf("  Índice %d: %s\n", i, fruta)

		// IMPORTANTE: Modificar 'fruta' NO modifica el array original
		// porque 'fruta' es una copia del valor, no una referencia
		fruta = "Modificado" // Esto NO afecta al array original
	}
	fmt.Printf("  Array después (no cambia): %v\n", frutas)

	// Método 3: Iteración con range, modificando el array original
	fmt.Println("\n1.3. Iteración con range modificando el array original:")
	for i := range frutas {
		// Aquí solo obtenemos el índice
		// Para modificar el array, usamos el índice directamente
		frutas[i] = strings.ToUpper(frutas[i])
	}
	fmt.Printf("  Array en mayúsculas: %v\n", frutas)

	// Método 4: Iteración con range, ignorando el índice
	fmt.Println("\n1.4. Iteración solo por valores (ignorando índice):")
	contador := 1
	for _, fruta := range frutas {
		// _ ignora el índice, fruta = valor
		fmt.Printf("  Fruta %d: %s\n", contador, fruta)
		contador++
	}

	// ==========================================
	// EJEMPLO 2: ITERACIÓN DE ARRAY CON FILTROS
	// ==========================================
	fmt.Println("\n2. ITERACIÓN CON FILTROS")

	// Array de números
	numeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Array de números: %v\n", numeros)

	// Ejemplo 2.1: Filtrado durante la iteración
	fmt.Println("\n2.1. Filtrado de números pares:")
	for i := range numeros {
		// Saltamos los números impares usando continue
		if numeros[i]%2 != 0 {
			continue // Salta a la siguiente iteración
		}
		fmt.Printf("  Número par en posición %d: %d\n", i, numeros[i])
	}

	// Ejemplo 2.2: Búsqueda y salida anticipada
	fmt.Println("\n2.2. Búsqueda con break:")
	valorBuscado := 7
	for i, num := range numeros {
		if num == valorBuscado {
			fmt.Printf("  ¡Encontrado! %d está en la posición %d\n", valorBuscado, i)
			break // Sale del bucle una vez encontrado
		}
		fmt.Printf("  Revisando posición %d: %d\n", i, num)
	}

	// Ejemplo 2.3: Acumulación durante la iteración
	fmt.Println("\n2.3. Acumulación de valores:")
	suma := 0
	for _, num := range numeros {
		suma += num // Acumulamos cada número
		// También podríamos hacer otras operaciones aquí
	}
	fmt.Printf("  La suma de todos los números es: %d\n", suma)
	fmt.Printf("  El promedio es: %.2f\n", float64(suma)/float64(len(numeros)))

	// ==========================================
	// EJEMPLO 3: ITERACIÓN DE ARRAYS MULTIDIMENSIONALES
	// ==========================================
	fmt.Println("\n3. ITERACIÓN DE ARRAYS MULTIDIMENSIONALES")

	// Array bidimensional (matriz)
	matriz := [3][4]int{
		{1, 2, 3, 4},    // Primera fila
		{5, 6, 7, 8},    // Segunda fila
		{9, 10, 11, 12}, // Tercera fila
	}
	fmt.Println("Matriz original:")
	for i := 0; i < len(matriz); i++ {
		fmt.Printf("  %v\n", matriz[i])
	}

	// Ejemplo 3.1: Iteración completa de matriz
	fmt.Println("\n3.1. Iteración completa de matriz:")
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			// i = índice de fila, j = índice de columna
			fmt.Printf("  matriz[%d][%d] = %d\n", i, j, matriz[i][j])
		}
	}

	// Ejemplo 3.2: Iteración con range de matriz
	fmt.Println("\n3.2. Iteración con range de matriz:")
	for i, fila := range matriz {
		// i = índice de fila, fila = array completo de esa fila
		fmt.Printf("  Fila %d: %v\n", i, fila)

		// Iteramos sobre cada elemento de la fila
		for j, valor := range fila {
			// j = índice de columna, valor = elemento
			fmt.Printf("    matriz[%d][%d] = %d\n", i, j, valor)
		}
	}

	// Ejemplo 3.3: Operaciones en una matriz
	fmt.Println("\n3.3. Duplicando valores en la diagonal principal:")
	for i := range matriz {
		for j := range matriz[i] {
			// Solo modificamos elementos en la diagonal principal (i == j)
			if i == j {
				matriz[i][j] *= 2 // Duplicamos el valor
				fmt.Printf("  Duplicando matriz[%d][%d] a %d\n", i, j, matriz[i][j])
			}
		}
	}

	// Mostramos la matriz resultante
	fmt.Println("  Matriz resultante:")
	for _, fila := range matriz {
		fmt.Printf("  %v\n", fila)
	}

	// ==========================================
	// EJEMPLO 4: TÉCNICAS AVANZADAS DE ITERACIÓN
	// ==========================================
	fmt.Println("\n4. TÉCNICAS AVANZADAS DE ITERACIÓN")

	// Array de calificaciones de estudiantes
	calificaciones := [5]int{85, 92, 78, 95, 88}
	fmt.Printf("Calificaciones: %v\n", calificaciones)

	// Ejemplo 4.1: Encontrar valor máximo y mínimo
	fmt.Println("\n4.1. Encontrar máximo y mínimo:")
	max, min := calificaciones[0], calificaciones[0]
	maxPos, minPos := 0, 0

	for i, cal := range calificaciones {
		// Actualizamos el máximo si encontramos valor mayor
		if cal > max {
			max = cal
			maxPos = i
		}

		// Actualizamos el mínimo si encontramos valor menor
		if cal < min {
			min = cal
			minPos = i
		}
	}

	fmt.Printf("  Calificación máxima: %d (posición %d)\n", max, maxPos)
	fmt.Printf("  Calificación mínima: %d (posición %d)\n", min, minPos)

	// Ejemplo 4.2: Iterar en orden inverso
	fmt.Println("\n4.2. Iteración en orden inverso:")
	for i := len(calificaciones) - 1; i >= 0; i-- {
		// i comienza desde el último índice y va hasta 0
		fmt.Printf("  calificaciones[%d] = %d\n", i, calificaciones[i])
	}

	// Ejemplo 4.3: Iteración por rangos específicos
	fmt.Println("\n4.3. Iteración por rangos específicos:")

	// Sólo los primeros 3 elementos
	fmt.Println("  Solo los primeros 3 elementos:")
	for i := 0; i < 3 && i < len(calificaciones); i++ {
		fmt.Printf("    calificaciones[%d] = %d\n", i, calificaciones[i])
	}

	// Sólo los últimos 2 elementos
	fmt.Println("  Solo los últimos 2 elementos:")
	for i := len(calificaciones) - 2; i < len(calificaciones); i++ {
		fmt.Printf("    calificaciones[%d] = %d\n", i, calificaciones[i])
	}

	// ==========================================
	// EJEMPLO 5: ITERACIÓN DE ARRAYS DE ESTRUCTURAS
	// ==========================================
	fmt.Println("\n5. ITERACIÓN DE ARRAYS DE ESTRUCTURAS")

	// Definimos un tipo Estudiante
	type Estudiante struct {
		Nombre string
		Edad   int
		Notas  [3]float64 // Array de 3 calificaciones
	}

	// Array de estudiantes
	estudiantes := [3]Estudiante{
		{"Ana", 20, [3]float64{9.5, 8.7, 9.2}},
		{"Juan", 22, [3]float64{7.8, 8.9, 8.5}},
		{"María", 21, [3]float64{9.9, 9.8, 9.7}},
	}

	fmt.Println("Lista de estudiantes:")

	// Iteración básica
	for i, est := range estudiantes {
		// Calculamos el promedio de calificaciones
		var sumaNotas float64
		for _, nota := range est.Notas {
			sumaNotas += nota
		}
		promedio := sumaNotas / float64(len(est.Notas))

		fmt.Printf("  Estudiante %d:\n", i+1)
		fmt.Printf("    Nombre: %s\n", est.Nombre)
		fmt.Printf("    Edad: %d años\n", est.Edad)
		fmt.Printf("    Notas: %v\n", est.Notas)
		fmt.Printf("    Promedio: %.2f\n", promedio)
	}

	// Ejemplo 5.1: Modificar estructuras durante iteración
	fmt.Println("\n5.1. Modificar estructuras durante iteración:")

	// IMPORTANTE: range en un array crea copias, debemos usar índices
	// para modificar el array original
	for i := range estudiantes {
		// Incrementamos la edad de cada estudiante
		estudiantes[i].Edad++

		// Agregamos 0.5 puntos a todas las notas
		for j := range estudiantes[i].Notas {
			estudiantes[i].Notas[j] += 0.5
			// Aseguramos que no pase de 10
			if estudiantes[i].Notas[j] > 10.0 {
				estudiantes[i].Notas[j] = 10.0
			}
		}
	}

	fmt.Println("Después de modificar:")
	for i, est := range estudiantes {
		fmt.Printf("  Estudiante %d: %s, %d años, Notas: %v\n",
			i+1, est.Nombre, est.Edad, est.Notas)
	}

	// ==========================================
	// EJEMPLO 6: ALGORITMOS COMUNES CON ARRAYS
	// ==========================================
	fmt.Println("\n6. ALGORITMOS COMUNES CON ARRAYS")

	// Array para ordenar
	desordenado := [7]int{64, 25, 12, 22, 11, 9, 37}
	fmt.Printf("Array original: %v\n", desordenado)

	// Ejemplo 6.1: Ordenamiento Burbuja
	fmt.Println("\n6.1. Ordenamiento de burbuja:")
	ordenado := desordenado // Creamos copia para ordenar

	for i := 0; i < len(ordenado)-1; i++ {
		// En cada iteración, el elemento más grande "burbujea" hasta el final
		fmt.Printf("  Pasada %d:\n", i+1)

		// Recorremos desde 0 hasta len-i-1 (la parte final ya está ordenada)
		for j := 0; j < len(ordenado)-i-1; j++ {
			// Si el elemento actual es mayor que el siguiente, los intercambiamos
			if ordenado[j] > ordenado[j+1] {
				fmt.Printf("    Intercambiando %d y %d\n", ordenado[j], ordenado[j+1])
				ordenado[j], ordenado[j+1] = ordenado[j+1], ordenado[j]
			}
		}

		fmt.Printf("    Estado actual: %v\n", ordenado)
	}

	fmt.Printf("  Array ordenado: %v\n", ordenado)

	// Ejemplo 6.2: Búsqueda binaria (en un array ordenado)
	fmt.Println("\n6.2. Búsqueda binaria:")
	valorBuscado = 22
	encontrado := false
	inicio, fin := 0, len(ordenado)-1

	fmt.Printf("  Buscando %d en el array ordenado...\n", valorBuscado)

	for inicio <= fin {
		medio := (inicio + fin) / 2
		fmt.Printf("    Examinando rango [%d:%d], valor medio en pos %d: %d\n",
			inicio, fin, medio, ordenado[medio])

		if ordenado[medio] == valorBuscado {
			fmt.Printf("    ¡Encontrado %d en la posición %d!\n", valorBuscado, medio)
			encontrado = true
			break
		} else if ordenado[medio] < valorBuscado {
			inicio = medio + 1 // Buscar en la mitad derecha
			fmt.Printf("    El valor es mayor, ajustando rango a [%d:%d]\n", inicio, fin)
		} else {
			fin = medio - 1 // Buscar en la mitad izquierda
			fmt.Printf("    El valor es menor, ajustando rango a [%d:%d]\n", inicio, fin)
		}
	}

	if !encontrado {
		fmt.Printf("  %d no se encontró en el array\n", valorBuscado)
	}

	// Ejemplo 6.3: Rotación de array
	fmt.Println("\n6.3. Rotación de array:")
	original := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	rotado := original // Copia para rotar
	posiciones := 3

	fmt.Printf("  Array original: %v\n", original)
	fmt.Printf("  Rotando %d posiciones a la derecha...\n", posiciones)

	// Método simple: repetimos la rotación una posición, varias veces
	for p := 0; p < posiciones; p++ {
		ultimo := rotado[len(rotado)-1] // Guardamos el último elemento

		// Movemos cada elemento una posición a la derecha
		for i := len(rotado) - 1; i > 0; i-- {
			rotado[i] = rotado[i-1]
		}

		rotado[0] = ultimo // El último va al principio
		fmt.Printf("    Después de rotación %d: %v\n", p+1, rotado)
	}

	fmt.Printf("  Array rotado final: %v\n", rotado)
}
