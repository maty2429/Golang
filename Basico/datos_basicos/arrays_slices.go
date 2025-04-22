package main

import "fmt"

func main() {
	// ===== ARRAYS =====
	// Los arrays en Go tienen tamaño fijo definido al momento de su creación
	// y no puede cambiar durante la ejecución del programa

	// Declaración de un array de enteros de tamaño 5
	// Por defecto, todos los elementos se inicializan con el valor cero del tipo (0 para int)
	var numeros [5]int

	// Los índices en Go comienzan en 0, al igual que en muchos otros lenguajes
	numeros[0] = 10 // Asigna 10 al primer elemento (posición 0)
	numeros[1] = 20 // Asigna 20 al segundo elemento (posición 1)
	// Las posiciones 2, 3 y 4 mantienen su valor por defecto (0)

	fmt.Println("Array completo:", numeros)     // Imprime: [10 20 0 0 0]
	fmt.Println("Primer elemento:", numeros[0]) // Accede al primer elemento (10)

	// Declaración e inicialización en una sola línea usando la sintaxis de inicialización corta (:=)
	// Aquí creamos un array de strings con 7 elementos ya inicializados
	dias := [7]string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
	fmt.Println("Días de la semana:", dias)

	// ===== ITERACIÓN DE ARRAYS =====
	// La forma más común de recorrer un array en Go es usando la palabra clave 'range'
	// 'range' devuelve dos valores por cada elemento: el índice y el valor
	for i, dia := range dias {
		// i = índice (comienza en 0)
		// dia = valor en la posición i

		// El formato %d es para números enteros, %s para strings
		// Sumamos 1 a i para mostrar días del 1 al 7 en lugar de 0 a 6
		fmt.Printf("Día %d: %s\n", i+1, dia)
	}

	// ===== SLICES =====
	// Los slices son referencias a arrays, más flexibles y de tamaño dinámico

	// Primero definimos un array normal de 5 elementos
	array := [5]int{1, 2, 3, 4, 5}

	// Creamos un slice tomando una sección del array:
	// array[inicio:fin] - incluye el elemento en 'inicio' pero excluye el elemento en 'fin'
	slice := array[1:4]          // Toma elementos de índices 1, 2, 3 (NO incluye el 4)
	fmt.Println("Slice:", slice) // Imprime: [2 3 4]

	// IMPORTANTE: los slices NO COPIAN los datos, solo apuntan al array original
	// Si modificas el slice, se modifica el array subyacente y viceversa

	// ===== SLICES LITERALES =====
	// También puedes crear slices directamente, sin declarar un array primero
	// Go creará un array subyacente automáticamente
	numeros2 := []int{10, 20, 30}           // Nota: sin tamaño entre [], esto lo hace un slice, no un array
	fmt.Println("Slice inicial:", numeros2) // Imprime: [10 20 30]

	// ===== APPEND EN SLICES =====
	// La función 'append' es clave - agrega elementos a un slice existente
	// Si no hay suficiente capacidad, Go creará un nuevo array subyacente más grande
	numeros2 = append(numeros2, 40, 50)               // Agrega 40 y 50 al final del slice
	fmt.Println("Slice después de append:", numeros2) // Imprime: [10 20 30 40 50]

	// ===== ITERACIÓN DE SLICES =====
	// Iterar sobre un slice es igual que iterar sobre un array
	for i, num := range numeros2 {
		// i = índice, num = valor en la posición i
		fmt.Printf("Índice %d: %d\n", i, num)
	}

	// ===== LONGITUD Y CAPACIDAD =====
	// len(): devuelve la cantidad de elementos en el slice (longitud)
	fmt.Println("Longitud del slice:", len(numeros2)) // 5 elementos

	// cap(): devuelve la capacidad máxima del slice antes de necesitar realocar memoria
	// La capacidad suele ser mayor o igual a la longitud
	fmt.Println("Capacidad del slice:", cap(numeros2))
	// La capacidad puede ser mayor que la longitud cuando Go ha reservado espacio extra
	// para futuras operaciones append, mejorando el rendimiento
}
