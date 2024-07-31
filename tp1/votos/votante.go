package votos

type TipoVoto int

const (
	PRESIDENTE TipoVoto = iota
	GOBERNADOR
	INTENDENTE
)

const (
	CANT_VOTACION = INTENDENTE + 1
	LISTA_IMPUGNA = 0
)

// Voto tiene guardada la información de un voto emitido, por cada tipo de voto posible.
// Por ejemplo, en la posición GOBERNADOR, tendrá guardada la alternativa a Gobernador.
// Si vale 0, es un voto en blanco.
// Si Impugnado es 'true', entonces no hay que considerar ninguna de las alterantivas señaladas.

type Voto struct {
	VotoPorTipo [CANT_VOTACION]int
	Impugnado   bool
}

// Votante modela un votante en nuestro sistema de votación
type Votante interface {

	//LeerDNI Nos da el DNI del votante
	LeerDNI() int

	//Votar asenta la alternativa elegida en el tipo de voto indicado. En caso que el votante ya hubiera terminado
	//anteriormente de votar, devolverá el error correspondiente. Sino, nil.
	Votar(tipo TipoVoto, alternativa int) error

	//Deshacer deshace la última operación realizada. Se tiene que poder deshacer hasta el estado inicial del voto
	//(equivalente a un voto completamente en blanco). En caso que efectivamente haya habido alguna acción para
	//deshacer, devolverá nil. En caso de no haber acción par adeshacer, devolverá el error correspondiente.
	//También puede devolver error en caso que el votante ya hubiera terminado antes su proceso de votación.
	Deshacer() error

	//FinVoto termina el proceso de votación para este votante. En caso que el votante ya hubiera terminado
	//anteriormente con el proceso de votación, devolverá el error correspondiente. Sino, el voto en el estado final
	//obtenido de las diferentes aplicaciones de Votar y Deshacer.
	FinVoto() (Voto, error)
}
