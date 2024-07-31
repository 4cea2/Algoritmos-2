package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
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

/*------EJ 12--------
Implementar una función func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) , que elimine los elementos encolados para los cuales la función filtro devuelve false.
Aquellos elementos que no son eliminados deben permanecer en el mismo orden en el que estaban antes de invocar a la función.
No es necesario destruir los elementos que sí fueron eliminados.
Se pueden utilizar las estructuras auxiliares que se consideren necesarias y no está permitido acceder a la estructura interna de la cola (es una función).
¿Cuál es el orden del algoritmo implementado?
*/

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
