package diccionario

import (
	"fmt"
	"math"
	TDAlista "tdas/lista"
)

const FACTOR_CARGA_AUMENTO float32 = 3
const FACTOR_AGRANDAR int = 2
const FACTOR_CARGA_REDUCCION float32 = 0.5
const FACTOR_ACHICAR int = 2
const TAMANOTABLAINICIAL int = 100

type campo[K comparable, V any] struct {
	clave K
	dato  V
}

type hashImplementacion[K comparable, V any] struct {
	tablaHash         []TDAlista.Lista[campo[K, V]]
	cantidadElementos int
	tamanoTabla       int
}

type iterhashImplementacion[K comparable, V any] struct {
	iterLista   TDAlista.IteradorLista[campo[K, V]]
	diccionario hashImplementacion[K, V]
	posDic      int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccionario := new(hashImplementacion[K, V])
	diccionario.tamanoTabla = TAMANOTABLAINICIAL
	diccionario.tablaHash = make([]TDAlista.Lista[campo[K, V]], diccionario.tamanoTabla)
	for idx, _ := range diccionario.tablaHash {
		diccionario.tablaHash[idx] = TDAlista.CrearListaEnlazada[campo[K, V]]()
	}
	diccionario.cantidadElementos = 0
	return diccionario
}

func (diccionario *hashImplementacion[K, V]) Guardar(clave K, dato V) {
	diccionario.redimensionarSiSeRequiere()
	iter := diccionario.buscarCampo(clave)
	posHash := funcionHashing(clave, diccionario.tamanoTabla)
	campoNuevo := campo[K, V]{clave, dato}
	if iter == nil {
		diccionario.tablaHash[posHash].InsertarUltimo(campoNuevo)
		diccionario.cantidadElementos++
		return
	}
	iter.Borrar()
	iter.Insertar(campoNuevo)
	return
}

// Busco la clave en el diccionario, en caso de que lo encuentre, devuelve el campo asociado (iterador).
// Caso contrario, devuelve nil.
func (diccionario *hashImplementacion[K, V]) buscarCampo(clave K) TDAlista.IteradorLista[campo[K, V]] {
	posHash := funcionHashing(clave, diccionario.tamanoTabla)
	listaActual := diccionario.tablaHash[posHash]
	iter := listaActual.Iterador()
	for iter.HaySiguiente() {
		campoActual := iter.VerActual()
		if clave == campoActual.clave {
			return iter
		}
		iter.Siguiente()
	}
	return nil
}

func (diccionario *hashImplementacion[K, V]) Pertenece(clave K) bool {
	iter := diccionario.buscarCampo(clave)
	if iter == nil {
		return false
	}
	return true
}

func (diccionario *hashImplementacion[K, V]) Obtener(clave K) V {
	iter := diccionario.buscarCampo(clave)
	if iter == nil {
		panic("La clave no pertenece al diccionario")
	}
	campoActual := iter.VerActual()
	return campoActual.dato
}

func (diccionario *hashImplementacion[K, V]) Borrar(clave K) V {
	iter := diccionario.buscarCampo(clave)
	if iter == nil {
		panic("La clave no pertenece al diccionario")
	}
	campoBorrado := iter.Borrar()
	diccionario.cantidadElementos--
	diccionario.redimensionarSiSeRequiere()
	return campoBorrado.dato
}

func (diccionario *hashImplementacion[K, V]) Cantidad() int {
	return diccionario.cantidadElementos
}

func (diccionario *hashImplementacion[K, V]) redimensionarSiSeRequiere() {
	factorDeCarga := float32(diccionario.cantidadElementos) / float32(diccionario.tamanoTabla)
	var nuevoTamanoTabla int
	if factorDeCarga > FACTOR_CARGA_AUMENTO {
		nuevoTamanoTabla = diccionario.tamanoTabla * FACTOR_AGRANDAR
	} else if (factorDeCarga < FACTOR_CARGA_REDUCCION) && (diccionario.tamanoTabla > TAMANOTABLAINICIAL) {
		nuevoTamanoTabla = (diccionario.tamanoTabla) / FACTOR_ACHICAR
	} else {
		return
	}
	tablaAnterior := diccionario.tablaHash
	nuevaTablaHash := make([]TDAlista.Lista[campo[K, V]], nuevoTamanoTabla)
	for idx, _ := range nuevaTablaHash {
		nuevaTablaHash[idx] = TDAlista.CrearListaEnlazada[campo[K, V]]()
	}
	for _, lista_i := range tablaAnterior {
		iter := lista_i.Iterador()
		for iter.HaySiguiente() {
			campoActual := iter.Borrar()
			posHash := funcionHashing(campoActual.clave, nuevoTamanoTabla)
			nuevaTablaHash[posHash].InsertarUltimo(campoActual)
		}
	}
	diccionario.tablaHash = nuevaTablaHash
	diccionario.tamanoTabla = nuevoTamanoTabla
}

func (diccionario *hashImplementacion[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	cortar := false
	for _, lista_i := range diccionario.tablaHash {
		lista_i.Iterar(func(par campo[K, V]) bool {
			if cortar {
				return false
			}
			if !visitar(par.clave, par.dato) {
				cortar = true
				return false
			}
			return true
		})

		if cortar {
			return
		}
	}
}

func (diccionario *hashImplementacion[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(iterhashImplementacion[K, V])
	iterador.diccionario = *diccionario
	posInicial := iterador.buscarSiguienteLista(0)
	if posInicial != -1 {
		iterador.posDic = posInicial
	} else {
		iterador.posDic = 0
	}
	iterador.iterLista = (iterador.diccionario.tablaHash[iterador.posDic]).Iterador()
	return iterador
}

// Busco la siguiente lista no vacia a partir de pos, en caso de que lo encuentre, devuelve la posicion de la lista nueva encontrada.
// Caso contrario, devuelve una posicion negativa.
func (iterador *iterhashImplementacion[K, V]) buscarSiguienteLista(pos int) int {
	for i := pos; i < iterador.diccionario.tamanoTabla; i++ {
		if !(iterador.diccionario.tablaHash[i].EstaVacia()) {
			return i
		}
	}
	return -1
}

func (iterador *iterhashImplementacion[K, V]) HaySiguiente() bool {
	return (iterador.iterLista.HaySiguiente())
}

func (iterador *iterhashImplementacion[K, V]) VerActual() (K, V) {
	if !(iterador.HaySiguiente()) {
		panic("El iterador termino de iterar")
	}
	campoActual := iterador.iterLista.VerActual()
	return campoActual.clave, campoActual.dato
}

func (iterador *iterhashImplementacion[K, V]) Siguiente() {
	if !(iterador.HaySiguiente()) {
		panic("El iterador termino de iterar")
	}
	iterador.iterLista.Siguiente()
	if !(iterador.iterLista.HaySiguiente()) {
		posActual := iterador.buscarSiguienteLista(iterador.posDic + 1)
		if posActual == -1 {
			return
		}
		iterador.posDic = posActual
		iterador.iterLista = (iterador.diccionario.tablaHash[iterador.posDic].Iterador())
	}
}

// ----------FUNCION HASHING----------
// Funcion sacada de: https://golangprojectstructure.com/hash-functions-go-code/
func funcionHashing[K comparable](clave K, tamArreglo int) int {
	bytes := convertirABytes(clave)
	return _funcionHashing(string(bytes), tamArreglo)
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// djb2 hash
func _funcionHashing(key string, tamano int) int {
	var posFinal uint64 = 5381
	for i := 0; i < len(key); i++ {
		posFinal += uint64(key[i]) + posFinal + posFinal<<5
	}
	return int(math.Abs(float64(int(posFinal) % tamano)))
}
