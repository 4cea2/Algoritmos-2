package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func nodoCrear[T any](dato T, prox *nodoCola[T]) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	nodo.prox = prox
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return ((cola.primero == nil) && (cola.ultimo == nil))
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nodo := nodoCrear(elem, nil)
	if cola.EstaVacia() {
		cola.primero = nodo
	} else {
		cola.ultimo.prox = nodo
	}
	cola.ultimo = nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.primero.dato
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return dato
}

/*
Implementar la primitiva func (cola *colaEnlazada[T]) Multiprimeros(k int) []T que dada una cola y un número k, devuelva los primeros k elementos de la cola, en el mismo orden en
el que habrían salido de la cola. En caso que la cola tenga menos de k elementos. Si hay menos elementos que k en la cola, devolver un slice del tamaño de la cola.
Indicar y justificar el orden de ejecución del algoritmo.

REHACES PORQUE ME PEDIAN DE ATRAS PA DELANTE, NO DE ADELANTE PA ATRAS

const CAPACIDAD int = 100

func (cola *colaEnlazada[T]) Multiprimeros(k int) []T {
	if cola.EstaVacia() {
		return make([]T, 0)
	}
	var i int
	ElementosDesencolados := make([]T, CAPACIDAD)
	NodoIterador := cola.primero
	ElementosDesencolados[0] = NodoIterador.dato
	for i = 1; (i < k) && (NodoIterador.prox != nil); i++ {
		NodoIterador = NodoIterador.prox
		ElementosDesencolados[i] = NodoIterador.dato
	}
	ArregloInvertido := make([]T, i)
	for j := 0; j < i; j++ {
		ArregloInvertido[j] = ElementosDesencolados[i-j-1]
	}
	return ArregloInvertido
}
*/
