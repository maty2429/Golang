package main

import "fmt"

func main() {
	fmt.Println("===== TUTORIAL DE ARRAYS Y SLICES EN GO =====")

	// ==========================================
	// PARTE 1: ARRAYS BÁSICOS
	// ==========================================
	fmt.Println("\n--- ARRAYS BÁSICOS ---")

	// En Go, los ARRAYS tienen tamaño FIJO y se declaran así:
	// var nombreArray [tamaño]tipo

	// Ejemplo 1: Array de enteros de tamaño 5 (todos inicializados a 0)
	var numeros [5]int
	fmt.Printf("Array recién creado: %v\n", numeros) // Imprime: [0 0 0 0 0]

	// Asignación elemento por elemento
	numeros[0] = 10 // Primera posición (índice 0)
	numeros[1] = 20 // Segunda posición (índice 1)
	numeros[4] = 50 // Última posición (índice 4)

	// IMPORTANTE: Acceder a un índice fuera del rango causa error en tiempo de ejecución
	// numeros[5] = 60 // Esto daría error: índice fuera de rango

	fmt.Printf("Array después de asignaciones: %v\n", numeros) // [10 20 0 0 50]

	// Ejemplo 2: Declaración e inicialización directa
	dias := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
	fmt.Printf("Array de días: %v\n", dias)

	// Ejemplo 3: Array con tamaño determinado por los elementos iniciales
	// El operador ... deja que Go cuente los elementos automáticamente
	colores := [...]string{"Rojo", "Verde", "Azul"}
	fmt.Printf("Array de colores (tamaño %d): %v\n", len(colores), colores)

	// ==========================================
	// PARTE 2: ITERACIÓN DE ARRAYS
	// ==========================================
	fmt.Println("\n--- ITERACIÓN DE ARRAYS ---")

	// Método 1: Iteración clásica con índices (como en C, Java)
	fmt.Println("Iteración con índices:")
	for i := 0; i < len(dias); i++ {
		fmt.Printf("  Día %d: %s\n", i+1, dias[i])
	}

	// Método 2: Iteración con range (recomendado en Go)
	fmt.Println("\nIteración con range (índice y valor):")
	for i, dia := range dias {
		fmt.Printf("  Posición %d: %s\n", i, dia)
	}

	// Si solo necesitas el valor y no el índice, puedes usar _ (guion bajo)
	fmt.Println("\nIteración con range (solo valores):")
	for _, dia := range dias {
		fmt.Printf("  %s\n", dia)
	}

	// Si solo necesitas el índice, omite la segunda variable
	fmt.Println("\nIteración con range (solo índices):")
	for i := range dias {
		fmt.Printf("  Índice: %d, Valor: %s\n", i, dias[i])
	}

	// ==========================================
	// PARTE 3: ARRAYS MULTIDIMENSIONALES
	// ==========================================
	fmt.Println("\n--- ARRAYS MULTIDIMENSIONALES ---")

	// Un array de arrays - una matriz 3x3
	matriz := [3][3]int{
		{1, 2, 3}, // Primera fila
		{4, 5, 6}, // Segunda fila
		{7, 8, 9}, // Tercera fila
	}

	fmt.Println("Matriz 3x3:")
	for fila := 0; fila < 3; fila++ {
		for columna := 0; columna < 3; columna++ {
			fmt.Printf("%d ", matriz[fila][columna])
		}
		fmt.Println() // Nueva línea después de cada fila
	}

	// ==========================================
	// PARTE 4: INTRODUCCIÓN A SLICES
	// ==========================================
	fmt.Println("\n--- INTRODUCCIÓN A SLICES ---")

	// DIFERENCIA CLAVE:
	// - Arrays: tamaño fijo, valor (se copian al asignar o pasar a funciones)
	// - Slices: tamaño dinámico, referencia (apuntan al array subyacente)

	// Crear un array base
	arrayBase := [5]int{10, 20, 30, 40, 50}

	// Crear slices desde el array (formato: array[inicio:fin])
	// RECUERDA: incluye el índice inicio, pero excluye el índice fin
	sliceCompleto := arrayBase[:] // Todos los elementos
	sliceInicio := arrayBase[0:3] // Elementos 0,1,2
	sliceFinal := arrayBase[2:5]  // Elementos 2,3,4
	sliceMedio := arrayBase[1:4]  // Elementos 1,2,3

	fmt.Printf("Array base: %v\n", arrayBase)
	fmt.Printf("Slice completo: %v\n", sliceCompleto)
	fmt.Printf("Slice inicio: %v\n", sliceInicio)
	fmt.Printf("Slice final: %v\n", sliceFinal)
	fmt.Printf("Slice medio: %v\n", sliceMedio)

	// IMPORTANTE: Los slices apuntan al array original
	// Si modificas el slice, el array cambia, y viceversa
	sliceMedio[0] = 999 // Cambia el elemento 1 del array original
	fmt.Println("\nDespués de modificar el slice:")
	fmt.Printf("Array base modificado: %v\n", arrayBase)   // [10 999 30 40 50]
	fmt.Printf("Slice medio modificado: %v\n", sliceMedio) // [999 30 40]

	// ==========================================
	// PARTE 5: CREACIÓN DIRECTA DE SLICES
	// ==========================================
	fmt.Println("\n--- CREACIÓN DIRECTA DE SLICES ---")

	// Método 1: Slice literal (Go crea un array subyacente automáticamente)
	nombres := []string{"Ana", "Juan", "Pedro"} // Sin tamaño definido = slice
	fmt.Printf("Slice de nombres: %v (longitud: %d)\n", nombres, len(nombres))

	// Método 2: Usando make([]tipo, longitud, capacidad)
	// make crea un slice con un array subyacente del tamaño especificado
	// - tipo: el tipo de datos
	// - longitud: número de elementos iniciales (inicializados a valor cero)
	// - capacidad: tamaño del array subyacente (opcional, por defecto = longitud)
	numeros2 := make([]int, 5) // Longitud 5, capacidad 5
	fmt.Printf("Slice con make: %v (longitud: %d, capacidad: %d)\n",
		numeros2, len(numeros2), cap(numeros2))

	numeros3 := make([]int, 3, 10) // Longitud 3, capacidad 10
	fmt.Printf("Slice con capacidad: %v (longitud: %d, capacidad: %d)\n",
		numeros3, len(numeros3), cap(numeros3))

	// ==========================================
	// PARTE 6: MANIPULANDO SLICES
	// ==========================================
	fmt.Println("\n--- MANIPULANDO SLICES ---")

	// Slice inicial
	frutas := []string{"Manzana", "Naranja", "Plátano"}
	fmt.Printf("Slice inicial: %v (len: %d, cap: %d)\n",
		frutas, len(frutas), cap(frutas))

	// append: agrega elementos al final del slice
	// IMPORTANTE: puede crear un nuevo array subyacente si no hay capacidad
	frutas = append(frutas, "Pera")
	frutas = append(frutas, "Uva", "Mango") // Puedes agregar múltiples elementos

	fmt.Printf("Después de append: %v (len: %d, cap: %d)\n",
		frutas, len(frutas), cap(frutas))
	// Nota: la capacidad puede haber aumentado (normalmente se duplica)

	// ==========================================
	// PARTE 7: CASOS PRÁCTICOS
	// ==========================================
	fmt.Println("\n--- CASO PRÁCTICO: FILTRADO ---")

	// Ejemplo: Filtrar números pares de un slice
	todosNumeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var numerosPares []int // Slice vacío

	for _, num := range todosNumeros {
		if num%2 == 0 { // Si es par
			numerosPares = append(numerosPares, num)
		}
	}

	fmt.Printf("Números originales: %v\n", todosNumeros)
	fmt.Printf("Números pares: %v\n", numerosPares)

	fmt.Println("\n--- CASO PRÁCTICO: SLICES DE SLICES ---")

	// Un slice puede contener otros slices (como una matriz irregular)
	// Útil para datos como una lista de estudiantes con calificaciones
	estudiantes := [][]string{
		{"Ana", "9", "8", "10"},
		{"Juan", "7", "9", "8"},
		{"Pedro", "10", "10", "9"},
	}

	fmt.Println("Lista de estudiantes y calificaciones:")
	fmt.Println("Nombre\tMat\tFís\tQui")
	fmt.Println("------\t---\t---\t---")

	for _, estudiante := range estudiantes {
		// Cada estudiante es un slice de strings
		for _, dato := range estudiante {
			fmt.Printf("%s\t", dato)
		}
		fmt.Println()
	}

	// ==========================================
	// PARTE 8: DIFERENCIAS CLAVE ARRAY VS SLICE
	// ==========================================
	fmt.Println("\n--- DIFERENCIAS CLAVE ARRAY VS SLICE ---")

	fmt.Println("1. Declaración:")
	fmt.Println("   - Array: [tamaño]tipo{valores}")
	fmt.Println("   - Slice: []tipo{valores} o make([]tipo, longitud, capacidad)")

	fmt.Println("\n2. Tamaño:")
	fmt.Println("   - Array: Fijo, no puede cambiar")
	fmt.Println("   - Slice: Dinámico, puede crecer con append()")

	fmt.Println("\n3. Al asignar o pasar a funciones:")
	fmt.Println("   - Array: Se copia completamente (valor)")
	fmt.Println("   - Slice: Se pasa la referencia (cambios afectan al original)")

	// Demostración: Array vs Slice en asignación
	arrayDemo := [3]int{1, 2, 3}
	sliceDemo := []int{1, 2, 3}

	// Copia de array (crea un nuevo array independiente)
	arrayCopiado := arrayDemo
	arrayCopiado[0] = 999

	// Referencia de slice (apunta al mismo array subyacente)
	sliceCopiado := sliceDemo
	sliceCopiado[0] = 999

	fmt.Printf("\nArray original después de modificar copia: %v\n", arrayDemo) // [1 2 3]
	fmt.Printf("Array copiado modificado: %v\n", arrayCopiado)                 // [999 2 3]

	fmt.Printf("Slice original después de modificar copia: %v\n", sliceDemo) // [999 2 3]
	fmt.Printf("Slice copiado modificado: %v\n", sliceCopiado)               // [999 2 3]

	fmt.Println("\nEsto demuestra que los arrays se copian por valor,")
	fmt.Println("mientras que los slices se pasan por referencia.")
}
