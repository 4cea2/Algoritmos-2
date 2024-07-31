package diccionario

type Diccionario[K comparable, V any] interface {

	// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
	Guardar(clave K, dato V)

	// Pertenece determina si una clave ya se encuentra en el diccionario, o no
	Pertenece(clave K) bool

	// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
	// 'La clave no pertenece al diccionario'
	Obtener(clave K) V

	// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
	// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
	Borrar(clave K) V

	// Cantidad devuelve la cantidad de elementos dentro del diccionario
	Cantidad() int

	// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del
	// mismo
	Iterar(func(clave K, dato V) bool)

	// Iterador devuelve un IterDiccionario para este Diccionario
	Iterador() IterDiccionario[K, V]
}

type IterDiccionario[K comparable, V any] interface {

	// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
	// el iterador hay un elemento.
	HaySiguiente() bool

	// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
	// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
	VerActual() (K, V)

	// Siguiente si HaySiguiente avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe
	// entrar en pánico con mensaje 'El iterador termino de iterar'
	Siguiente()
}
