package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con la creacion de la pila")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "La pila esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestLIFO(t *testing.T) {
	t.Log("Hacemos pruebas para ver si se mantiene la invariante de la pila")
	pila := TDAPila.CrearPilaDinamica[string]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	pila.Apilar("Hola")
	require.EqualValues(t, "Hola", pila.VerTope())
	pila.Apilar("que")
	require.EqualValues(t, "que", pila.VerTope())
	pila.Apilar("tal!")
	require.EqualValues(t, "tal!", pila.Desapilar())
	require.EqualValues(t, "que", pila.Desapilar())
	require.EqualValues(t, "Hola", pila.Desapilar())
	TestPilaVacia(t)
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 1000; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for i := 999; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
	TestPilaVacia(t)
}

/*
[1, 3, 5, 10, 11, 5, 2, 20] ---> [1, 3, 11, 2]


const CAPACIDAD int = 100

func TestEjRecu(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(3)
	pila.Apilar(5)
	pila.Apilar(2)
	pila.Apilar(10)
	pila.Apilar(7)
	var ElementosDesapilados [CAPACIDAD]int
	var i int
	for i = 0; !pila.EstaVacia(); i++ {
		if (pila.VerTope() % 5) == 0 {
			pila.Desapilar()
			i--
			continue
		}
		ElementosDesapilados[i] = pila.Desapilar()
	}
	for j := 0; j < i; j++ {
		pila.Apilar(ElementosDesapilados[i-j-1])
		fmt.Println(pila.VerTope())
	}
}

¿Orden?
Las primitivas del TDA pila se ejecutan de manera constante, osea, O(1)
Al querer desapilar todos los elementos de la pila (n), tengo una complejidad temporal de O(n)
Cuando vuelvo a colocar los elementos que no son multiplos de 5 (k) en la pila, obtengo, O(k)
Complejidad temporal: O(n+k)
Si me paro en el peor de los caso, obtengo que k = n (todos los elementos desapilados son multiplos de 5)
Por ende, la complejidad temporal me terminaria dando: O(n).
*/

/* ----------------EJ 11---------------------------------
Implementar una función que ordene de manera ascendente una pila de enteros sin conocer su estructura interna y utilizando como estructura auxiliar sólo otra pila auxiliar.
Por ejemplo, la pila [ 4, 1, 5, 2, 3 ] debe quedar como [ 1, 2, 3, 4, 5 ] (siendo el último elemento el tope de la pila, en ambos casos).
Indicar y justificar el orden de la función.


const CAPACIDAD int = 5

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func TestEj11(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(4)
	pila.Apilar(1)
	pila.Apilar(5)
	pila.Apilar(2)
	pila.Apilar(3)
	Elementos := make([]int, CAPACIDAD)
	var i int
	for i = 0; !pila.EstaVacia(); i++ {
		Elementos[i] = pila.Desapilar()
	}
	Elementos = mergeSort(Elementos)
	fmt.Println(Elementos)
	for j := 0; j < i; j++ {
		pila.Apilar(Elementos[j])
	}
	require.EqualValues(t, 5, pila.Desapilar())
	require.EqualValues(t, 4, pila.Desapilar())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
}
*/

/* ------------------EJ 6---------------------------------
Dada una pila de enteros, escribir una función que determine si sus elementos están ordenados de manera ascendente.
Una pila de enteros está ordenada de manera ascendente si, en el sentido que va desde el tope de la pila hacia el resto de elementos, cada elemento es menor al elemento que le sigue.
La pila debe quedar en el mismo estado que al invocarse la función. Indicar y justificar el orden del algoritmo propuesto.

const CAPACIDAD int = 100

func ApilarNuevamente(pila Pila[int], elementos []int, cantidad int) {
	for i := cantidad; i >= 0; i-- {
		pila.Apilar(elementos[i])
	}
}

func EsPiramidal(pila Pila[int]) bool {
	if pila.EstaVacia() {
		return true
	}
	ElementosDesapilados := [CAPACIDAD]int
	ElementosDesapilados[0] = pila.Desapilar()
	var i int
	for i = 1; !pila.EstaVacia; i++ {
		ElementosDesapilados[i] = pila.Desapilar()
		if ElementosDesapilados[i] <= ElementosDesapilados[i-1] {
			ApilarNuevamente(pila, ElementosDesapilados, i)
			return false
		}
	}
	ApilarNuevamente(pila, ElementosDesapilados, i)
	return true
}

¿Orden?
La declaracion de las variables enteras y arreglos lo puedo tomar como O(1) (incluye los ifs).
Al entrar al for, lo que hago es desapilar la pila, que en el peor caso puede ser O(n) si tengo que desapilar todos (n: cantidad de elementos desapilados).
A su vez, en el for tengo la funcion ApilarNuevamente, que apilaria los n elementos en caso de detectar que la pila no es piramidal.
Igualmente despues del for llamo a la funcion ApilarNuevamente, en el caso que sea piramidal.
Como en cada ciclo for no ejecuto a la funcion ApilarNuevamente (al menos cuando no es piramidal), puedo asumir que tengo una complejidad de O(n) + O(n) = O(n).



*/
