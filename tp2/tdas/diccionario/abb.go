package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterABBdict[K comparable, V any] struct {
	pilaNodos  TDAPila.Pila[*nodoAbb[K, V]]
	abb        *abb[K, V]
	esConRango bool
	desde      *K
	hasta      *K
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	arbol := new(abb[K, V])
	arbol.cmp = funcion_cmp
	return arbol

}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	if abb.raiz == nil && abb.cantidad == 0 {
		abb.raiz = crearNodoABB(clave, dato)
		abb.cantidad++
		return
	}
	abb._guardar(&abb.raiz, clave, dato)
}

/*
	 estuve mejorando... esto deberia funcionar pero nop :'(
		func (abb *abb[K, V]) Guardar(clave K, dato V) {
			if abb.raiz == nil && abb.cantidad == 0 {
				abb.raiz = crearNodoABB(clave, dato)
				abb.cantidad++
				return
			}
			nodo := abb._buscarNodo(&abb.raiz, clave)
			if nodo != nil {
				(*nodo).clave = clave
				(*nodo).dato = dato
			} else {
				*nodo = crearNodoABB(clave, dato)
				abb.cantidad++
			}
		}
*/
func (abb *abb[K, V]) _guardar(padre **nodoAbb[K, V], clave K, dato V) {
	if *padre == nil {
		*padre = crearNodoABB(clave, dato)
		abb.cantidad++
		return
	}
	if abb.cmp(clave, (*padre).clave) == 0 {
		(*padre).clave = clave
		(*padre).dato = dato
		return
	}
	if abb.cmp(clave, (*padre).clave) < 0 {
		abb._guardar(&(*padre).izquierdo, clave, dato)
	} else {
		abb._guardar(&(*padre).derecho, clave, dato)
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	if abb.raiz == nil && abb.cantidad == 0 {
		return false
	}
	return abb._buscarNodo(&abb.raiz, clave) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb._buscarNodo(&abb.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	//busco el la flechita al nodo a borrar
	pFlechita := abb._buscarNodo(&abb.raiz, clave)
	if pFlechita == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := (*pFlechita).dato
	//caso con cero hijos
	if (*pFlechita).izquierdo == nil && (*pFlechita).derecho == nil {
		*pFlechita = nil
		abb.cantidad--
		return dato
	}
	//caso con un hijo:
	if (*pFlechita).izquierdo == nil {
		*pFlechita = (*pFlechita).derecho
	} else if (*pFlechita).derecho == nil {
		*pFlechita = (*pFlechita).izquierdo
	} else {
		//caso con dos hijos
		reemplazante := (*pFlechita).izquierdo
		actual := reemplazante.derecho
		if actual == nil {
			reemplazante.derecho = (*pFlechita).derecho
			*pFlechita = reemplazante
			abb.cantidad--
			return dato
		}
		for actual.derecho != nil {
			reemplazante = actual
			actual = reemplazante.derecho
		}
		reemplazante.derecho = actual.izquierdo
		(*pFlechita).clave = actual.clave
		(*pFlechita).dato = actual.dato
	}
	abb.cantidad--
	return dato
}

// Retorna el ->padre.izquierdo o ->padre.derecho segun si el hijo con clave K es
// izquierdo o derecho respectivamente y si el padre es el nil retorna nil
func (abb *abb[K, V]) _buscarNodo(padre **nodoAbb[K, V], clave K) **nodoAbb[K, V] {
	if *padre == nil {
		return nil
	}
	if abb.cmp(clave, (**padre).clave) < 0 {
		return abb._buscarNodo(&(**padre).izquierdo, clave)
	} else if abb.cmp(clave, (**padre).clave) > 0 {
		return abb._buscarNodo(&(**padre).derecho, clave)
	}
	return padre
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

//
// -------------------------------ITERADORES--------------------------------------
//

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if !abb._iterarRango(abb.raiz, desde, hasta, visitar) {
		return
	}
}
func (abb *abb[K, V]) _iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if (desde == nil || abb.cmp(nodo.clave, *desde) > 0) && !abb._iterarRango(nodo.izquierdo, desde, hasta, visitar) {
		return false
	}
	if (desde == nil || abb.cmp(nodo.clave, *desde) >= 0) && (hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if (hasta == nil || abb.cmp(nodo.clave, *hasta) < 0) && !abb._iterarRango(nodo.derecho, desde, hasta, visitar) {
		return false
	}
	return true
}

// -

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := abb.crearIterador()
	iter.esConRango = false
	abb.raiz.apilarHijosIzquierdos(iter)
	return iter
}

// Apila en el iterador externo todos los hijos izquierdos del nodo.
func (nodo *nodoAbb[K, V]) apilarHijosIzquierdos(iter *iterABBdict[K, V]) {
	if nodo == nil {
		return
	}
	if iter.esConRango && iter.desde != nil && !(iter.abb.cmp(nodo.clave, *(iter.desde)) >= 0) {
		nodo.derecho.apilarHijosIzquierdos(iter)
		return
	}
	iter.pilaNodos.Apilar(nodo)
	nodo.izquierdo.apilarHijosIzquierdos(iter)
}

func (iter *iterABBdict[K, V]) HaySiguiente() bool {
	if iter.pilaNodos.EstaVacia() {
		return false
	}
	nodoActual := iter.pilaNodos.VerTope()
	if iter.esConRango && iter.hasta != nil && !(iter.abb.cmp(*(iter.hasta), nodoActual.clave) >= 0) {
		return false
	}
	return true
}

func (iter *iterABBdict[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoActual := iter.pilaNodos.VerTope()
	if nodoActual.izquierdo == nil && nodoActual.derecho == nil {
		iter.pilaNodos.Desapilar()
	} else {
		nodoDerecho := nodoActual.derecho
		iter.pilaNodos.Desapilar()
		nodoDerecho.apilarHijosIzquierdos(iter)
	}
}

func (iter *iterABBdict[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoActual := iter.pilaNodos.VerTope()
	return nodoActual.clave, nodoActual.dato
}

// Crea el iterador externo.
func (abb *abb[K, V]) crearIterador() *iterABBdict[K, V] {
	iter := new(iterABBdict[K, V])
	iter.pilaNodos = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.abb = abb
	return iter
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := abb.crearIterador()
	iter.esConRango = true
	iter.desde = desde
	iter.hasta = hasta
	abb.raiz.apilarHijosIzquierdos(iter)
	return iter
}
