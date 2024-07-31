package cola_prioridad_test

import (
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const TAMANOVOLUMEN int = 10000

func cmpMaximoInt(k1 int, k2 int) int {
	return k1 - k2
}

func cmpMinimoInt(k1 int, k2 int) int {
	return k2 - k1
}

func cmpMaximoString(k1 string, k2 string) int {
	return len(k1) - len(k2)
}

func cmpMinimoString(k1 string, k2 string) int {
	return len(k2) - len(k1)
}

func cmpStringAscii(k1 string, k2 string) int {
	return int(k1[0]) - int(k2[0])
}

func TestHeapVacio(t *testing.T) {
	t.Log("Creo un heap vacio")
	heap := TDAHeap.CrearHeap[int](cmpMaximoInt)
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapArrVacio(t *testing.T) {
	t.Log("Creo un heap vacio con un arreglo.")
	var arreglo []int
	heap := TDAHeap.CrearHeapArr[int](arreglo, cmpMaximoInt)
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapUnElemento(t *testing.T) {
	t.Log("Creo un heap con 1 elemento")
	heap := TDAHeap.CrearHeap[int](cmpMaximoInt)
	heap.Encolar(10)
	require.Equal(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.Equal(t, 10, heap.VerMax())
}

func TestHeapArrUnElemento(t *testing.T) {
	t.Log("Creo un heap con un arreglo de 1 elemento, o con un arreglo vacio y le agrego un elemento")
	//Primera forma
	arregloConUnElemento := []int{10}
	heap := TDAHeap.CrearHeapArr[int](arregloConUnElemento, cmpMaximoInt)
	require.Equal(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.Equal(t, 10, heap.VerMax())
	//Segunda forma
	arregloVacio := []int{}
	heap = TDAHeap.CrearHeapArr[int](arregloVacio, cmpMaximoInt)
	heap.Encolar(10)
	require.Equal(t, 1, heap.Cantidad())
	require.False(t, heap.EstaVacia())
	require.Equal(t, 10, heap.VerMax())
}

func TestHeapMinimo(t *testing.T) {
	t.Log("Creo un heap de minimo de tipo int y string")
	//Int
	arregloInicialInt := []int{5, 9, 10, 12, 11, 2}
	arregloFinalInts := []int{2, 5, 9, 10, 11, 12}

	heapInt := TDAHeap.CrearHeap[int](cmpMinimoInt)
	posMinimo := 0
	heapInt.Encolar(arregloInicialInt[0])
	//Encolo y verifico el minimo en todo momento.
	for i := 1; i < len(arregloInicialInt); i++ {
		heapInt.Encolar(arregloInicialInt[i])
		if cmpMinimoInt(arregloInicialInt[posMinimo], arregloInicialInt[i]) < 0 {
			posMinimo = i
		}
		require.Equal(t, arregloInicialInt[posMinimo], heapInt.VerMax())
	}
	//Desencolo y verifico el minimo en todo momento.
	for i := 0; i < len(arregloFinalInts)-1; i++ {
		require.Equal(t, arregloFinalInts[i], heapInt.Desencolar())
		require.Equal(t, arregloFinalInts[i+1], heapInt.VerMax())
	}

	//String
	arregloInicialStrings := []string{"hola", "rosita", "te", "quiero", "mucho", "!"}
	arregloFinalStrings := []string{"!", "te", "hola", "mucho", "rosita", "quiero"}

	heapString := TDAHeap.CrearHeap[string](cmpMinimoString)
	posMinimo = 0
	heapString.Encolar(arregloInicialStrings[0])
	//Encolo y verifico el minimo en todo momento.
	for i := 1; i < len(arregloInicialStrings); i++ {
		heapString.Encolar(arregloInicialStrings[i])
		if cmpMinimoString(arregloInicialStrings[posMinimo], arregloInicialStrings[i]) < 0 {
			posMinimo = i
		}
		require.Equal(t, arregloInicialStrings[posMinimo], heapString.VerMax())
	}
	//Desencolo y verifico el minimo en todo momento.
	for i := 0; i < len(arregloFinalStrings)-1; i++ {
		require.Equal(t, arregloFinalStrings[i], heapString.Desencolar())
		require.Equal(t, arregloFinalStrings[i+1], heapString.VerMax())
	}
}

func TestHeapMaximo(t *testing.T) {
	t.Log("Creo un heap de maximo de tipo int y string")
	//Int
	arregloInicialInt := []int{5, 9, 10, 12, 11, 2}
	arregloFinalInts := []int{12, 11, 10, 9, 5, 2}

	heapInt := TDAHeap.CrearHeap[int](cmpMaximoInt)
	posMaximo := 0
	heapInt.Encolar(arregloInicialInt[0])
	//Encolo y verifico el maximo en todo momento.
	for i := 1; i < len(arregloInicialInt); i++ {
		heapInt.Encolar(arregloInicialInt[i])
		if cmpMaximoInt(arregloInicialInt[posMaximo], arregloInicialInt[i]) < 0 {
			posMaximo = i
		}
		require.Equal(t, arregloInicialInt[posMaximo], heapInt.VerMax())
	}
	//Desencolo y verifico el maximo en todo momento.
	for i := 0; i < len(arregloFinalInts)-1; i++ {
		require.Equal(t, arregloFinalInts[i], heapInt.Desencolar())
		require.Equal(t, arregloFinalInts[i+1], heapInt.VerMax())
	}

	//String
	arregloInicialStrings := []string{"hola", "rosita", "te", "quiero", "mucho", "!"}
	arregloFinalStrings := []string{"rosita", "quiero", "mucho", "hola", "te", "!"}

	heapString := TDAHeap.CrearHeap[string](cmpMaximoString)
	posMaximo = 0
	heapString.Encolar(arregloInicialStrings[0])
	//Encolo y verifico el maximo en todo momento.
	for i := 1; i < len(arregloInicialStrings); i++ {
		heapString.Encolar(arregloInicialStrings[i])
		if cmpMaximoString(arregloInicialStrings[posMaximo], arregloInicialStrings[i]) < 0 {
			posMaximo = i
		}
		require.Equal(t, arregloInicialStrings[posMaximo], heapString.VerMax())
	}
	//Desencolo y verifico el maximo en todo momento.
	for i := 0; i < len(arregloFinalStrings)-1; i++ {
		require.Equal(t, arregloFinalStrings[i], heapString.Desencolar())
		require.Equal(t, arregloFinalStrings[i+1], heapString.VerMax())
	}
}

func TestHeapArrMaximo(t *testing.T) {
	t.Log("Creo un heap maximo de tipo int y string con un arreglo de las 2 formas posibles")

	arregloVacioInt := []int{}
	arregloInicialInt := []int{5, 10, 11, 9, 12, 2}
	arregloFinalInts := []int{12, 11, 10, 9, 5, 2}

	arregloVacioStrings := []string{}
	arregloInicialStrings := []string{"mucho", "!", "te", "hola", "quiero", "rosita"}
	arregloFinalStrings := []string{"quiero", "rosita", "mucho", "hola", "te", "!"}
	//-----Primera forma-----

	//Int
	heapInt := TDAHeap.CrearHeapArr[int](arregloInicialInt, cmpMaximoInt)
	for i := 0; i < len(arregloFinalInts)-1; i++ {
		require.Equal(t, arregloFinalInts[i], heapInt.Desencolar())
		require.Equal(t, arregloFinalInts[i+1], heapInt.VerMax())
	}

	//String
	heapString := TDAHeap.CrearHeapArr[string](arregloInicialStrings, cmpMaximoString)
	for i := 0; i < len(arregloFinalStrings)-1; i++ {
		require.Equal(t, arregloFinalStrings[i], heapString.Desencolar())
		require.Equal(t, arregloFinalStrings[i+1], heapString.VerMax())
	}
	//-----Segunda forma-----
	posMaximo := 0
	//Int
	heapInt = TDAHeap.CrearHeapArr[int](arregloVacioInt, cmpMaximoInt)
	heapInt.Encolar(arregloInicialInt[0])
	for i := 1; i < len(arregloInicialInt); i++ {
		heapInt.Encolar(arregloInicialInt[i])
		if cmpMaximoInt(arregloInicialInt[posMaximo], arregloInicialInt[i]) < 0 {
			posMaximo = i
		}
		require.Equal(t, arregloInicialInt[posMaximo], heapInt.VerMax())
	}
	for i := 0; i < len(arregloFinalInts)-1; i++ {
		require.Equal(t, arregloFinalInts[i], heapInt.Desencolar())
		require.Equal(t, arregloFinalInts[i+1], heapInt.VerMax())
	}

	//String
	heapString = TDAHeap.CrearHeapArr[string](arregloVacioStrings, cmpMaximoString)
	heapString.Encolar(arregloInicialStrings[0])
	posMaximo = 0
	for i := 1; i < len(arregloInicialStrings); i++ {
		heapString.Encolar(arregloInicialStrings[i])
		if cmpMaximoString(arregloInicialStrings[posMaximo], arregloInicialStrings[i]) < 0 {
			posMaximo = i
		}
		require.Equal(t, arregloInicialStrings[posMaximo], heapString.VerMax())
	}
	for i := 0; i < len(arregloFinalStrings)-1; i++ {
		require.Equal(t, arregloFinalStrings[i], heapString.Desencolar())
		require.Equal(t, arregloFinalStrings[i+1], heapString.VerMax())
	}
}

func TestHeapArrMinimo(t *testing.T) {
	t.Log("Creo un heap minimo de tipo int y string con un arreglo de las 2 formas posibles")

	arregloVacioInt := []int{}
	arregloInicialInt := []int{5, 10, 11, 9, 12, 2}
	arregloFinalInts := []int{2, 5, 9, 10, 11, 12}

	arregloVacioStrings := []string{}
	arregloInicialStrings := []string{"mucho", "!", "te", "hola", "quiero", "rosita"}
	arregloFinalStrings := []string{"!", "te", "hola", "mucho", "rosita", "quiero"}
	//-----Primera forma-----

	//Int
	heapInt := TDAHeap.CrearHeapArr[int](arregloInicialInt, cmpMinimoInt)
	for i := 0; i < len(arregloFinalInts)-1; i++ {
		require.Equal(t, arregloFinalInts[i], heapInt.Desencolar())
		require.Equal(t, arregloFinalInts[i+1], heapInt.VerMax())
	}

	//String
	heapString := TDAHeap.CrearHeapArr[string](arregloInicialStrings, cmpMinimoString)
	for i := 0; i < len(arregloFinalStrings)-1; i++ {
		require.Equal(t, arregloFinalStrings[i], heapString.Desencolar())
		require.Equal(t, arregloFinalStrings[i+1], heapString.VerMax())
	}

	//-----Segunda forma-----
	posMinimo := 0
	//Int
	heapInt = TDAHeap.CrearHeapArr[int](arregloVacioInt, cmpMinimoInt)
	heapInt.Encolar(arregloInicialInt[0])
	for i := 1; i < len(arregloInicialInt); i++ {
		heapInt.Encolar(arregloInicialInt[i])
		if cmpMinimoInt(arregloInicialInt[posMinimo], arregloInicialInt[i]) < 0 {
			posMinimo = i
		}
		require.Equal(t, arregloInicialInt[posMinimo], heapInt.VerMax())
	}
	for i := 0; i < len(arregloFinalInts)-1; i++ {
		require.Equal(t, arregloFinalInts[i], heapInt.Desencolar())
		require.Equal(t, arregloFinalInts[i+1], heapInt.VerMax())
	}

	//String
	heapString = TDAHeap.CrearHeapArr[string](arregloVacioStrings, cmpMinimoString)
	heapString.Encolar(arregloInicialStrings[0])
	posMinimo = 0
	//Encolo y verifico el minimo en todo momento.
	for i := 1; i < len(arregloInicialStrings); i++ {
		heapString.Encolar(arregloInicialStrings[i])
		if cmpMinimoString(arregloInicialStrings[posMinimo], arregloInicialStrings[i]) < 0 {
			posMinimo = i
		}
		require.Equal(t, arregloInicialStrings[posMinimo], heapString.VerMax())
	}
	//Desencolo y verifico el minimo en todo momento
	for i := 0; i < len(arregloFinalStrings)-1; i++ {
		require.Equal(t, arregloFinalStrings[i], heapString.Desencolar())
		require.Equal(t, arregloFinalStrings[i+1], heapString.VerMax())
	}
}

func TestHeapSort(t *testing.T) {
	t.Log("Test del metodo de ordenamiento heapsort con tipo int y string")
	// Arreglo desordenado
	arregloInt := []int{9, 2, 7, 4, 1, 5, 8, 3, 6}
	arregloString := []string{"B", "b", "D", "a", "c", "A", "C"}

	// Ordenar el arreglo utilizando HeapSort
	TDAHeap.HeapSort(arregloInt, cmpMaximoInt)
	TDAHeap.HeapSort(arregloString, cmpStringAscii)
	// Verificar que el arreglo esté ordenado
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, arregloInt, "El arreglo no está ordenado")
	require.Equal(t, []string{"A", "B", "C", "D", "a", "b", "c"}, arregloString, "malo malo")
}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Test de volumen del metodo de ordenamiento heapsort")
	arregloIntInicial := make([]int, TAMANOVOLUMEN)
	arregloIntFinal := make([]int, TAMANOVOLUMEN)

	for i := 0; i < TAMANOVOLUMEN; i++ {
		arregloIntInicial[i] = i
		arregloIntFinal[i] = TAMANOVOLUMEN - i - 1
	}
	TDAHeap.HeapSort(arregloIntInicial, cmpMinimoInt)
	require.Equal(t, arregloIntFinal, arregloIntInicial)
}

func TestHeapVolumen(t *testing.T) {
	t.Log("Test de volumen de un heap minimo")
	heap := TDAHeap.CrearHeap[int](cmpMinimoInt)
	arregloIntFinal := make([]int, TAMANOVOLUMEN)

	//Encolo y verifico siempre el maximo.
	heap.Encolar(0)
	for i := 1; i <= TAMANOVOLUMEN/2; i++ {
		heap.Encolar(i)
		heap.Encolar((-1) * i)
		require.Equal(t, (-1)*i, heap.VerMax())
	}
	//Asigno al arreglo final como deberia estar el heap
	for i := -TAMANOVOLUMEN / 2; i < TAMANOVOLUMEN/2; i++ {
		arregloIntFinal[TAMANOVOLUMEN/2+i] = i
	}
	//Desencolo y verifico siempre el maximo.
	for i := 0; i < len(arregloIntFinal)-1; i++ {
		require.Equal(t, arregloIntFinal[i], heap.Desencolar())
		require.Equal(t, arregloIntFinal[i+1], heap.VerMax())
	}
}

func TestHeapArrVolumen(t *testing.T) {
	t.Log("Test de volumen de un heap maximo con un arreglo")
	arregloIntInicial := make([]int, TAMANOVOLUMEN)
	arregloIntFinal := make([]int, TAMANOVOLUMEN)
	//Asigno al arreglo inicial
	for i := (-TAMANOVOLUMEN / 2) + 1; i <= TAMANOVOLUMEN/2; i++ {
		arregloIntInicial[i+TAMANOVOLUMEN/2-1] = i
	}
	//Asigno al arreglo final como deberia estar el heap
	for i := (TAMANOVOLUMEN / 2); i > -TAMANOVOLUMEN/2; i-- {
		arregloIntFinal[TAMANOVOLUMEN/2-i] = i
	}
	heap := TDAHeap.CrearHeapArr[int](arregloIntInicial, cmpMaximoInt)
	//Desencolo y verifico siempre el maximo.
	for i := 0; i < len(arregloIntFinal)-1; i++ {
		require.Equal(t, arregloIntFinal[i], heap.Desencolar())
		require.Equal(t, arregloIntFinal[i+1], heap.VerMax())
	}
}
