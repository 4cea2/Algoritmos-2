package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	CAPACIDADINICIAL = 5
	AGRANDAR         = 2
	REDUCIR          = 2
	MULTIPLICADOR    = 4
)

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDADINICIAL)
	return pila
}

func (pila *pilaDinamica[T]) redimensionarSiSeRequiere() {
	capacidadActual := cap(pila.datos)
	if pila.cantidad == capacidadActual {
		capacidadNueva := AGRANDAR * capacidadActual
		array := make([]T, capacidadNueva)
		copy(array, pila.datos)
		pila.datos = array
	} else if (pila.cantidad*MULTIPLICADOR <= capacidadActual) && (capacidadActual > CAPACIDADINICIAL) {
		capacidadNueva := capacidadActual / REDUCIR
		array := make([]T, capacidadNueva)
		copy(array, pila.datos)
		pila.datos = array
	}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	pila.redimensionarSiSeRequiere()
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	elementoDesapilado := pila.datos[pila.cantidad-1]
	pila.cantidad--
	pila.redimensionarSiSeRequiere()
	return elementoDesapilado
}
