package filaVotos

import votos "rerepolez/votos"

// SistemaDeVotacion modela un sistema de votacion.
type SistemaDeVotacion interface {

	//El votante ingresa a la fila con su respectivo dni.
	//Devolvera el error correspondiente si no se encuentra registrado el dni.
	//Caso contrario, nil.
	ComandoIngresar(dni int) error

	//El votante votara a su representante de cierto partido.
	//Devolvera el error correspondiente si no hay nadie en la fila para votar, o la alternativa sea invalida, o que quiera votar aunque ya haya finalizado su voto (fraudulento).
	//aunque ya haya finalizado su voto con anterioridad (fraudulento).
	//Caso contrario, nil.
	ComandoVotar(tipoVoto votos.TipoVoto, alternativa string) error

	//El votante deshace la ultima accion hecha, ya sea votar o quiere des-impugnarse.
	//Devolvera el error correspondiente si no hay nadie en la fila para deshacer su voto, o no haya voto a deshacer, o hay un votante en la fila que quiere
	//deshacer su voto, aunque ya haya finalizado su voto con anterioridad (fraudulento).
	//Caso contrario, nil.
	ComandoDeshacer() error

	//El votante finaliza su voto.
	//Devolvera el error correspondiente si no hay nadie en la fila para finalizar su voto, o ya haya finalizando su voto con anterioridad (fraudulento).
	//Caso contrario, nil.
	ComandoFinVotar() error

	//Finaliza el proceso de votacion del sistema.
	//Devuelve los votos finales en un arreglo y su respectiva cantidad.
	//Caso en que haya votantes aun en la fila esperando para votar, devolvera el error correspondiente.
	//Caso contrario, nil.
	ComandoFinalizarVotacion() ([]votos.Voto, int, error)
}
