package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	max := 0
	for i := 1; i < len(vector); i++ {
		if vector[max] < vector[i] {
			max = i
		}
	}
	return max
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	long1 := len(vector1)
	long2 := len(vector2)
	var son_iguales bool = false
	var v1_mayor_v2 bool = false
	var n int
	if long1 < long2 {
		n = long1
	} else if long1 > long2 {
		n = long2
		v1_mayor_v2 = true
	} else {
		n = long1
		son_iguales = true
	}
	for i := 0; i < n; i++ {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}
	if son_iguales {
		return 0
	}
	if v1_mayor_v2 {
		return 1
	}
	return -1
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := len(vector); i != 0; i-- {
		Swap(&vector[Maximo(vector[:i])], &vector[i-1])
	}
}

// Funcion auxiliar para la funcion Suma
func aux_Suma(vector []int, n int) int {
	if n == 0 {
		return vector[n]
	}
	return vector[n] + aux_Suma(vector, n-1)
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return aux_Suma(vector, len(vector)-1)
}

// Funcion auxiliar para la funcion EsCadenaCapicua
func es_palindromo(s string, inicio, fin int) bool {
	if inicio >= fin {
		return true
	}
	if s[inicio] != s[fin] {
		return false
	}
	return es_palindromo(s, inicio+1, fin-1)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return es_palindromo(cadena, 0, len(cadena)-1)
}
