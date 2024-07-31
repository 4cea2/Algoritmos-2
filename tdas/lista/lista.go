package lista

type IteradorLista[T any] interface {
	//Muestra el valor actual que se esta iterando.
	VerActual() T
	//Retorna true o false si hay elemento para itera o no, respectivamente.
	HaySiguiente() bool
	//Avanza el iterador al siguiente elemnto.
	Siguiente()
	//Inserta un nuevo elemento antes  del iterador, mueve el iterador en el elemento insertado.
	Insertar(T)
	//Borrar el elemento donde esta el iterador, el iterador avanza al siguiente.
	Borrar() T
}

type Lista[T any] interface {
	//Retorna true si la lista esta vacia.
	EstaVacia() bool
	//Inserta al principio de la lista.
	InsertarPrimero(T)
	//inserta al final de la lista.
	InsertarUltimo(T)
	//Borra el primer elemento de la lista.
	BorrarPrimero() T
	//Muestra el dato del primer elemento valido de la lista.
	VerPrimero() T
	//Muestra el dado del ultimo elemento valido de la lista.
	VerUltimo() T
	//Retorna el largo de la lista.
	Largo() int
	//Retorna el iterador interno, visitar se aplica a cada elemento iterado siempre que retorne true.
	Iterar(visitar func(T) bool)
	//Retorna un iterador externo, que permiten usar los metodos de iterador lista.
	Iterador() IteradorLista[T]
}
