package votos

// Partido modela un partido político, con sus alternativas para cada uno de los tipos de votaciones
type Partido interface {

	//VotadoPara indica que este Partido ha recibido un voto para el TipoVoto indicado. Felicitaciones!
	VotadoPara(tipo TipoVoto)

	//ObtenerResultado permite obtener el resultado de este Partido para el TipoVoto indicado. El formato será el
	//conveniente para ser mostrado.
	ObtenerResultado(tipo TipoVoto) string
}
