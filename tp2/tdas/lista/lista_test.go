package lista_test

import (
	TDAlista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const MAX_ITER = 1000

func TestListaVacia(t *testing.T) {
	t.Log("Test de lista vacia ")
	lista := TDAlista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestListaInsertar(t *testing.T) {
	t.Log("Test de lista al insertar ")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	TestListaVacia(t)
}

func TestStress(t *testing.T) {
	t.Log("Tests stress")
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 1; i <= MAX_ITER; i++ {
		lista.InsertarPrimero(i)
	}
	for i := MAX_ITER; i > 0; i-- {
		aux := lista.BorrarPrimero()
		require.EqualValues(t, i, aux, "Se esperaba  %d", i)
	}
	TestListaVacia(t)
	for i := 1; i <= MAX_ITER; i++ {
		lista.InsertarUltimo(i)
	}
	for i := 1; i <= MAX_ITER; i++ {
		require.EqualValues(t, i, lista.BorrarPrimero(), "Se esperaba  %d", i)
	}
	TestListaVacia(t)
}

func TestString(t *testing.T) {
	t.Log("Test con una lista de string")
	listaString := TDAlista.CrearListaEnlazada[string]()
	alphabet := [26]string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	for _, e := range alphabet {
		listaString.InsertarUltimo(e)
	}
	for _, e := range alphabet {
		require.EqualValues(t, e, listaString.BorrarPrimero(), "Se esperaba obtener %s", e)
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaString.BorrarPrimero() })
}

func TestPtrString(t *testing.T) {
	t.Log("Test con una lista de punteros a string")
	listaPString := TDAlista.CrearListaEnlazada[*string]()
	alphabet := [26]string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
	for i := range alphabet {
		listaPString.InsertarUltimo(&alphabet[i])
	}
	for _, e := range alphabet {
		ptrS := listaPString.BorrarPrimero()
		require.EqualValues(t, e, *ptrS, "Se esperaba obtener %s", e)
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaPString.BorrarPrimero() })
}

func TestIteradorFinal(t *testing.T) {
	t.Log("Tests iterador con print")
	var suma int
	iter := 10
	sumaEsperada := iter * (iter + 1) / 2

	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 1; i <= iter; i++ {
		lista.InsertarPrimero(i)
	}

	lista.Iterar(func(v int) bool {
		require.EqualValues(t, iter, v)
		suma += v
		iter--
		return true

	})
	//valido que la funcion se haya aplicado:
	require.Equal(t, sumaEsperada, suma)

}

func TestIterHastaX(t *testing.T) {
	t.Log("Tests iterador con condicion de corte por false")
	var suma int
	var x int //x < iter
	iter := 10
	sumaTot := iter * (iter + 1) / 2

	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 1; i <= iter; i++ {
		lista.InsertarPrimero(i)
	}
	lista.Iterar(func(v int) bool {
		suma += v
		require.Equal(t, iter, v)
		if v == x {
			require.EqualValues(t, x, v)
		}
		iter--
		return v > x

	})
	n := x - 1                          //1+2+3+4
	sumaEsperada := sumaTot - n*(n+1)/2 //10+9+8+7+4+5
	//valido que la funcion se haya aplicado:
	require.Equal(t, sumaEsperada, suma)
}

func TestIterInsertar(t *testing.T) {
	t.Log("Test de insertar con el iterador")
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	require.Equal(t, 1, iter.VerActual())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	iter.Insertar(2)
	require.Equal(t, 2, iter.VerActual())
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
	iter.Insertar(3)
	require.Equal(t, 3, iter.VerActual())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}

func TestIterSiguiente(t *testing.T) {
	t.Log("Test de siguiente del iterador")
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	iter.Insertar(2)
	require.Equal(t, 2, iter.VerActual())
}

func TestIterListaInsertar(t *testing.T) {
	t.Log("Tests de insertar primero y ultimo")
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	require.Equal(t, 1, iter.VerActual())
	lista.InsertarUltimo(2)
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual())
}

func TestIterInsertarFinal(t *testing.T) {
	t.Log("Tests de insertar ultimo con el iterador")
	countIter := 9
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	for i := 2; i <= countIter; i++ {
		require.Equal(t, i-1, lista.VerPrimero())
		iter.Insertar(i)
	}
	require.Equal(t, 9, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	for ; iter.HaySiguiente(); iter.Siguiente() {
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	iter.Insertar(10)
	require.Equal(t, 10, lista.VerUltimo())
	require.Equal(t, 10, lista.Largo())
}

func TestIterInsertarEnElMedio(t *testing.T) {
	t.Log("Tests de insertar en el medio con iterador")
	countIter := 9
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	for i := 2; i <= countIter; i++ {
		require.Equal(t, i-1, lista.VerPrimero())
		iter.Insertar(i)
	}
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 4, iter.VerActual())
	iter.Insertar(100)
	require.Equal(t, 100, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 4, iter.VerActual())
}

func TestIterRemover(t *testing.T) {
	t.Log("Tests de Remover en varias posiciones con el iterador")
	countIter := 9
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(1)
	for i := 2; i <= countIter; i++ {
		require.Equal(t, i-1, lista.VerPrimero())
		iter.Insertar(i)
	}

	require.Equal(t, 9, iter.Borrar(), "deberia recibir un 9")

	require.Equal(t, 8, lista.Largo())
	require.Equal(t, 8, lista.VerPrimero())
	require.Equal(t, 8, iter.VerActual())
	iter.Siguiente()
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 5, iter.VerActual())
	iter.Borrar()
	require.Equal(t, 4, iter.VerActual())
	require.Equal(t, 7, lista.Largo())

	for iter.HaySiguiente() {
		iter.Siguiente()
		if iter.VerActual() == 1 {
			break
		}
	}

	iter.Borrar()
	require.Equal(t, 6, lista.Largo())

	require.Equal(t, 2, lista.VerUltimo())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.Equal(t, 8, lista.VerPrimero())

}

func TestIterarInicioFin(t *testing.T) {
	iter := 1000
	lista := TDAlista.CrearListaEnlazada[int]()
	for i := 1; i <= iter; i++ {
		lista.InsertarUltimo(i)
	}
	var i int = 1
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		require.Equal(t, i, iter.VerActual())
		i++
	}
	require.Equal(t, iter, lista.Largo())
}

/*
func TestInsertarAlFinalVerPrincio(t *testing.T) {
	max := 20
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for i := 1; i <= max; i++ {
		iter.Insertar(i)
	}
	require.Equal(t, max, lista.VerPrimero(), "Error: se modifico las referencias iniciales")
	iter.Siguiente()
	require.Equal(t, max-1, iter.VerActual(), "Error: se modifico las referencias iniciales")
	require.Equal(t, 1, lista.VerUltimo(), "Error: se modifico las referencias iniciales")
	//require.Equal(t, 2, iter.VerActual(), "Error: se modifico la segunda referencia")

}
*/

// Prueba borrar en una lista con un único elemento, y la lista queda en un estado correcto para
// insertar luego con primitivas de lista

func TestIteradorBorraListaConUnicoElemento(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(123123)
	iter := lista.Iterador()
	require.Equal(t, 123123, iter.Borrar())

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())

	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
}

// Insetarmos al inicio con iterador, y la lista queda en estado válido.
// Iteramos un paso mas para volver a llegar al final y luego inserto al final con primitiva de lista,
// itero con otro iterador y borramos todos los elementos con primitiva de lista para ver que esté todo en orden

func TestIterInsertarAlFinalListaNoVacia(t *testing.T) {
	var i int = 1
	lista := TDAlista.CrearListaEnlazada[int]()
	iter := lista.Iterador() //[*0,]:largo:0
	iter.Insertar(i)         //[*1,]:largo:1
	iter.Insertar(i + 1)     //[*1,1]:largo:2

	require.Equal(t, 2, lista.Largo())
	iter.Siguiente() //[2,*1,]:largo:2
	iter.Siguiente() //[2,1,*]:largo:2
	require.False(t, iter.HaySiguiente())

	lista.InsertarUltimo(i + 2) //(2,1,3):largo:3

	otroIter := lista.Iterador()                 //[*2, 1, 3]:largo:3
	require.Equal(t, i+1, lista.BorrarPrimero()) //(1,3):largo:2
	otroIter.Siguiente()                         //[*1,3]:largo:2
	require.Equal(t, i, lista.BorrarPrimero())   //(3):largo:1
	otroIter.Siguiente()                         //[*3]:largo:1
	require.Equal(t, i+2, lista.BorrarPrimero()) //():largo:0

	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.False(t, iter.HaySiguiente())

}
