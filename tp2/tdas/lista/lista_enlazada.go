package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	nodoIter *nodoLista[T]
	nodoAnte *nodoLista[T]
	lista    *listaEnlazada[T]
}

func nodoCrear[T any](dato T, prox *nodoLista[T]) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	nodo.siguiente = prox
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil && lista.ultimo == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(t T) {
	primerNodo := nodoCrear(t, lista.primero)
	if lista.EstaVacia() {
		lista.ultimo = primerNodo
	}
	lista.primero = primerNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(t T) {
	if lista.EstaVacia() {
		lista.InsertarPrimero(t)
		return
	}
	lista.ultimo.siguiente = nodoCrear(t, nil)
	lista.ultimo = lista.ultimo.siguiente
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	ptr := lista.primero
	for ptr != nil {
		if !visitar(ptr.dato) {
			return
		}
		ptr = ptr.siguiente
	}

}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista.primero, lista.primero, lista}
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	if iter.nodoIter == nil {
		return false
	}
	return true
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.nodoIter.dato
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if iter.HaySiguiente() {
		iter.nodoAnte = iter.nodoIter
		iter.nodoIter = iter.nodoIter.siguiente
		return
	}
	panic("El iterador termino de iterar")
}

// El elemento insertado va a tomar la posicion del elemento al que se apunta.
// Luego de una insercion, el iterador va a apuntar al nuevo elemento.
func (iter *iterListaEnlazada[T]) Insertar(t T) {
	iter.lista.largo++
	nodoNuevo := nodoCrear(t, iter.nodoIter)
	var entroEnUnIf bool

	if !iter.HaySiguiente() {
		entroEnUnIf = iter.lista.EstaVacia()
		iter.lista.ultimo = nodoNuevo
	}
	if iter.lista.primero == iter.nodoIter {
		entroEnUnIf = true
	}

	if entroEnUnIf {
		iter.lista.primero = nodoNuevo
		iter.nodoAnte = nodoNuevo
	} else {
		iter.nodoAnte.siguiente = nodoNuevo
	}
	iter.nodoIter = nodoNuevo
}

func (iter *iterListaEnlazada[T]) Borrar() T {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iter.nodoIter.dato
	if iter.nodoIter == iter.lista.ultimo {
		if iter.lista.largo == 1 {
			iter.lista.ultimo = iter.lista.ultimo.siguiente
		} else {
			iter.lista.ultimo = iter.nodoAnte
		}
	}
	if iter.nodoIter == iter.lista.primero {
		iter.lista.primero = iter.lista.primero.siguiente
		iter.nodoAnte = iter.nodoAnte.siguiente
	} else {
		iter.nodoAnte.siguiente = iter.nodoIter.siguiente
	}
	iter.nodoIter = iter.nodoIter.siguiente
	iter.lista.largo--
	return dato
}

/*

































func (i *iterListaEnlazada[T]) Insertar(t T) {
	i.lista.largo++

	nodoNuevo := nodoCrear(t, i.nodoIter)
	if i.nodoIter == nil {
		if i.lista.EstaVacia() {
			i.lista.primero = nodoNuevo
		}
		i.lista.ultimo = nodoNuevo
	} else if i.nodoIter == i.lista.primero {
		i.lista.primero = nodoNuevo
	} else {
		i.nodoAnte.siguiente = nodoNuevo
	}
	i.nodoAnte = nodoNuevo
	i.nodoIter = nodoNuevo

}





	func (i *iterListaEnlazada[T]) Insertar(t T) {
		i.lista.largo++
		nodoNuevo := nodoCrear(t, i.nodoIter)
		var entroEnUnIf bool =

		if i.lista.EstaVacia() {
			entroEnUnIf = true
			nodoNuevo.siguiente = nil
			i.lista.ultimo = nodoNuevo
		}
		if i.lista.primero == i.nodoIter {
			entroEnUnIf = true
		}
		if entroEnUnIf {
			i.lista.primero = nodoNuevo
			i.nodoAnte = nodoNuevo
			i.nodoIter = nodoNuevo
			return
		}
		i.nodoAnte.siguiente = nodoNuevo
		if !i.HaySiguiente() {
			i.lista.ultimo = nodoNuevo
		}
		i.nodoIter = nodoNuevo
	}








func (i *iterListaEnlazada[T]) Borrar() T {

	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	if i.lista.largo == 1 {
		i.nodoAnte = nil
		i.nodoIter = nil
		return i.lista.BorrarPrimero()
	}
	i.lista.largo--
	nodoDato := i.nodoIter.dato
	//cuando estoy en el ultimo elemento lista.ultimo->iter quiero: lista.ultimo->ante
	if i.nodoIter.siguiente == nil {
		i.lista.ultimo = i.nodoAnte
	}
	//estoy iterando el primer nodo, lo voy a borrar, debo cambiar Lista.Primero
	if i.lista.primero == i.nodoIter {
		i.lista.primero = i.nodoIter.siguiente
	}
	//caso general//el primero y ultimo no se deberian modificar.
	i.nodoAnte.siguiente = i.nodoIter.siguiente
	i.nodoIter = i.nodoIter.siguiente
	return nodoDato
}
*/
