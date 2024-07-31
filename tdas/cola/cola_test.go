package cola_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con la creacion de la cola")
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "La cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestFIFO(t *testing.T) {
	t.Log("Hacemos pruebas para ver si se mantiene la invariante de la cola")
	cola := TDACola.CrearColaEnlazada[string]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	cola.Encolar("Hola")
	require.EqualValues(t, "Hola", cola.VerPrimero())
	cola.Encolar("que")
	cola.Encolar("tal!")
	require.EqualValues(t, "Hola", cola.Desencolar())
	require.EqualValues(t, "que", cola.Desencolar())
	require.EqualValues(t, "tal!", cola.Desencolar())
	TestColaVacia(t)
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 1000; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 1000; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
	TestColaVacia(t)
}

const CAPACIDADINICIAL int = 5

func partir[T any](org TDACola.Cola[T], k int) []TDACola.Cola[T] {
	if org.EstaVacia() {
		return make([]TDACola.Cola[T], 0)
	}
	capacidadActual := CAPACIDADINICIAL
	colasFinales := make([]TDACola.Cola[T], CAPACIDADINICIAL)
	var colasUsadas int
	yaTermine := false
	for colasUsadas = 0; !yaTermine; colasUsadas++ {
		if colasUsadas == capacidadActual {
			capacidadActual *= 2
			arr := make([]TDACola.Cola[T], capacidadActual)
			copy(arr, colasFinales)
			colasFinales = arr
		}
		colaAuxiliar := TDACola.CrearColaEnlazada[T]()
		for j := 0; j < k; j++ {
			if org.EstaVacia() {
				yaTermine = true
				break
			}
			colaAuxiliar.Encolar(org.Desencolar())
		}
		colasFinales[colasUsadas] = colaAuxiliar
	}
	return colasFinales[:colasUsadas]
}

const DIAOK int = 20

func todoOkElDia(n int) bool {
	if n <= DIAOK {
		return true
	} else {
		return false
	}
}

func buscarDiaFalla(diasTotales int) int {
	return _buscarDia(diasTotales, 0, diasTotales-1)
}

func _buscarDia(dia, inicio, fin int) int {
	if fin-inicio <= 0 {
		return dia
	}
	medio := (inicio + fin) / 2
	if todoOkElDia(medio) {
		return _buscarDia(medio, medio+1, fin)
	} else {
		return _buscarDia(medio, inicio, medio)
	}

}

// 1er Parcial
func Test1erParcialito(t *testing.T) {
	org := TDACola.CrearColaEnlazada[int]()
	elementosEncolar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(elementosEncolar); i++ {
		org.Encolar(elementosEncolar[i])
	}
	k := 2
	colasFinales := partir[int](org, k)
	for i := 0; i < len(colasFinales); i++ {
		colaActual := colasFinales[i]
		fmt.Printf("Cola %d:", i)
		for !colaActual.EstaVacia() {
			fmt.Printf("%d, ", colaActual.Desencolar())
		}
		fmt.Printf("\n")
	}
	require.True(t, true, true)
	diaFinal := buscarDiaFalla(60)
	require.Equal(t, 21, diaFinal)
}

/*------EJ 12--------
Implementar una función func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) , que elimine los elementos encolados para los cuales la función filtro devuelve false.
Aquellos elementos que no son eliminados deben permanecer en el mismo orden en el que estaban antes de invocar a la función.
No es necesario destruir los elementos que sí fueron eliminados.
Se pueden utilizar las estructuras auxiliares que se consideren necesarias y no está permitido acceder a la estructura interna de la cola (es una función).
¿Cuál es el orden del algoritmo implementado?
*/

/*

func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) {
	if cola.EstaVacia() {
		return
	}
	colaAuxiliar := CrearColaEnlazada[K]()
	for !cola.EstaVacia() {
		if !filtro(cola.VerPrimero()) {
			cola.Desencolar()
			continue
		}
		colaAuxiliar.Encolar(cola.Desencolar())
	}
	cola = colaAuxiliar
}

*/
