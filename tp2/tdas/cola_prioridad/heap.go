package cola_prioridad

type fcmpHeap[K any] func(K, K) int

type heap[T comparable] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

const (
	CAPACIDADINICIAL = 5
	AGRANDAR         = 2
	REDUCIR          = 2
	MULTIPLICADOR    = 4
)

func CrearHeap[T comparable](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.cmp = funcion_cmp
	heap.datos = make([]T, CAPACIDADINICIAL)
	return heap
}

func CrearHeapArr[T comparable](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	cantidad := len(arreglo)
	arregloCopia := make([]T, cantidad)
	copy(arregloCopia, arreglo)
	heapify[T](arregloCopia, cantidad, funcion_cmp)
	heap.cmp = funcion_cmp
	heap.cantidad = cantidad
	heap.datos = arregloCopia
	if heap.cantidad < CAPACIDADINICIAL {
		nuevoArreglo := make([]T, CAPACIDADINICIAL)
		copy(nuevoArreglo, heap.datos)
		heap.datos = nuevoArreglo
	}
	return heap
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.Cantidad() == 0
}

func (heap *heap[T]) Encolar(dato T) {
	esNecesarioRedimensionar, capNueva := heap.hayQueRedimensionar()
	if esNecesarioRedimensionar {
		heap.redimensionar(capNueva)
	}
	heap.datos[heap.cantidad] = dato
	upheap[T](heap.datos, heap.cantidad, heap.cmp)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	datoDesencolado := heap.datos[0]
	heap.cantidad--
	swap(&heap.datos[0], &heap.datos[heap.cantidad])
	downheap[T](heap.datos, 0, heap.cantidad, heap.cmp)
	esNecesarioRedimensionar, capNueva := heap.hayQueRedimensionar()
	if esNecesarioRedimensionar {
		heap.redimensionar(capNueva)
	}
	return datoDesencolado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

// Metodo de ordenamiento comparativo
func HeapSort[T comparable](elementos []T, funcion_cmp func(T, T) int) {
	cantidad := len(elementos)
	heapify(elementos, cantidad, funcion_cmp)
	for i := cantidad - 1; i >= 0; i-- {
		swap(&elementos[0], &elementos[i])
		downheap(elementos, 0, i, funcion_cmp)
	}
}

// Hace dowheap del ultimo al primer elemento.
func heapify[T comparable](datos []T, cantidad int, cmp fcmpHeap[T]) {
	for i := cantidad - 1; i >= 0; i-- {
		downheap(datos, i, cantidad, cmp)
	}
}

func upheap[T comparable](datos []T, posHijo int, cmp func(T, T) int) {
	if posHijo == 0 {
		return
	}
	posPadre := posicionPadre(posHijo)
	prioriPadre := cmp(datos[posHijo], datos[posPadre])
	if prioriPadre > 0 {
		swap(&datos[posHijo], &datos[posPadre])
		upheap(datos, posPadre, cmp)
	}
}

func downheap[T comparable](datos []T, posPadre int, cantidad int, cmp func(T, T) int) {
	posHijoIzq := posicionHijoIzquierdo(posPadre)
	posHijoDer := posicionHijoDerecho(posPadre)
	if posHijoIzq >= cantidad {
		return
	}
	posPrioriMax := posicionPrioridadMaxima[T](datos, posPadre, posHijoIzq, posHijoDer, cantidad, cmp)
	if posPrioriMax != posPadre {
		swap(&datos[posPadre], &datos[posPrioriMax])
		downheap(datos, posPrioriMax, cantidad, cmp)
	}
}

func posicionPadre(posHijo int) int {
	return (posHijo - 1) / 2
}

func posicionHijoIzquierdo(posPadre int) int {
	return 2*posPadre + 1
}

func posicionHijoDerecho(posPadre int) int {
	return 2*posPadre + 2
}

func swap[T comparable](x *T, y *T) {
	*x, *y = *y, *x
}

// Devuelve la posicion del hijo que tenga mas prioridad con respecto a la posicion del padre.
// En caso de que ninguno de sus hijos tenga prioridad, devuelve la posicion del padre.
func posicionPrioridadMaxima[T comparable](datos []T, posPadre, posHijIzq, posHijDer, cantidad int, cmp func(T, T) int) int {
	prioriHijoIzq := cmp(datos[posHijIzq], datos[posPadre])
	if posHijDer >= cantidad {
		if prioriHijoIzq > 0 {
			return posHijIzq
		}
		return posPadre
	}

	prioriHijoDer := cmp(datos[posHijDer], datos[posPadre])
	if (prioriHijoDer < 0) && (prioriHijoIzq < 0) {
		return posPadre
	}
	prioriFinal := cmp(datos[posHijIzq], datos[posHijDer])
	if prioriFinal > 0 {
		return posHijIzq
	} else {
		return posHijDer
	}
}

// Verifica si hay que redimensionar el heap. En caso de hacerlo, devuelve true y la capacidad a redimensionar.
// Caso contrario, devuelve false y la capacidad en 0.
func (heap *heap[T]) hayQueRedimensionar() (bool, int) {
	capacidadActual := cap(heap.datos)
	var capacidadNueva int
	if heap.cantidad == capacidadActual {
		capacidadNueva = capacidadActual * AGRANDAR
		return true, capacidadNueva
	} else if (heap.cantidad*MULTIPLICADOR <= capacidadActual) && (capacidadActual > CAPACIDADINICIAL) {
		capacidadNueva = capacidadActual / REDUCIR
		return true, capacidadNueva
	} else {
		capacidadNueva = 0
		return false, capacidadNueva
	}
}

// Redimensiona el heap con la nueva capacidad.
func (heap *heap[T]) redimensionar(capNueva int) {
	array := make([]T, capNueva)
	copy(array, heap.datos)
	heap.datos = array
}
