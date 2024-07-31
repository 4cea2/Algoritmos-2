package filaVotos

import (
	Error "rerepolez/errores"
	votante "rerepolez/votos"
	"strconv"
	TDACola "tdas/cola"
)

type sistemaDeVotacionImplementacion struct {
	padronesRegistrados []int             //Padrones registrados que pueden votar.
	filaVotantes        []votante.Votante //Votantes registrados que pueden votar.
	padronesIngresados  TDACola.Cola[int] //Ingresos de padrones que van a votar.
	votosFinales        []votante.Voto    //Votos finales de los que ingresaron.
	tamanoVotos         int               //Cantidad de votos finales.
	tamanoPartidos      int               //Cantidad de partidos presentados.
}

const VOTOSINICIALES int = 5

func CrearSistemaVotacion(padrones []int, tamanoPadrones int, tamanoDePartidos int) SistemaDeVotacion {
	sistema := new(sistemaDeVotacionImplementacion)
	sistema.padronesRegistrados = padrones
	sistema.filaVotantes = make([]votante.Votante, tamanoPadrones)
	for i := 0; i < tamanoPadrones; i++ {
		sistema.filaVotantes[i] = votante.CrearVotante(padrones[i])
	}
	sistema.padronesIngresados = TDACola.CrearColaEnlazada[int]()
	sistema.votosFinales = make([]votante.Voto, VOTOSINICIALES)
	sistema.tamanoPartidos = tamanoDePartidos
	return sistema
}

func (sist *sistemaDeVotacionImplementacion) ComandoIngresar(NumeroDNI int) error {
	estaEnPadron, posPadron := dniEnPadron(NumeroDNI, sist)
	var errPadron error = nil
	if !estaEnPadron {
		errPadron = Error.DNIFueraPadron{}
		return errPadron
	}
	sist.padronesIngresados.Encolar(posPadron)
	return errPadron
}

func (sist *sistemaDeVotacionImplementacion) ComandoVotar(tipoVoto votante.TipoVoto, alternativa string) error {
	if sist.padronesIngresados.EstaVacia() {
		return Error.FilaVacia{}
	}
	NumeroLista, errConversion := strconv.Atoi(alternativa)
	if (errConversion != nil) || NumeroLista >= sist.tamanoPartidos {
		return Error.ErrorAlternativaInvalida{}
	}
	posPadron := sist.padronesIngresados.VerPrimero()
	votanteActual := sist.filaVotantes[posPadron]
	err := votanteActual.Votar(tipoVoto, NumeroLista)
	if err != nil {
		sist.padronesIngresados.Desencolar()
		return err
	}
	sist.filaVotantes[posPadron] = votanteActual
	return nil
}

func (sist *sistemaDeVotacionImplementacion) ComandoDeshacer() error {
	if sist.padronesIngresados.EstaVacia() {
		return Error.FilaVacia{}
	}
	posPadron := sist.padronesIngresados.VerPrimero()
	votanteActual := sist.filaVotantes[posPadron]
	err := votanteActual.Deshacer()
	if err != nil {
		if (err.Error() == Error.ErrorVotanteFraudulento{Dni: votanteActual.LeerDNI()}.Error()) {
			sist.padronesIngresados.Desencolar()
		}
		return err
	}
	sist.filaVotantes[posPadron] = votanteActual
	return nil
}

func (sist *sistemaDeVotacionImplementacion) ComandoFinVotar() error {
	if sist.padronesIngresados.EstaVacia() {
		return Error.FilaVacia{}
	}
	posPadron := sist.padronesIngresados.Desencolar()
	votanteActual := sist.filaVotantes[posPadron]
	votoFinal, err := votanteActual.FinVoto()
	if err != nil {
		return err
	}
	sist.filaVotantes[posPadron] = votanteActual
	capacidadActual := cap(sist.votosFinales)
	if sist.tamanoVotos == capacidadActual {
		capacidadNuevo := 2 * capacidadActual
		nuevosVotosFinales := make([]votante.Voto, capacidadNuevo)
		copy(nuevosVotosFinales, sist.votosFinales)
		sist.votosFinales = nuevosVotosFinales
	}
	sist.votosFinales[sist.tamanoVotos] = votoFinal
	sist.tamanoVotos++
	return nil
}

func (sist *sistemaDeVotacionImplementacion) ComandoFinalizarVotacion() ([]votante.Voto, int, error) {
	var errFinal error = nil
	if !(sist.padronesIngresados.EstaVacia()) {
		errFinal = Error.ErrorCiudadanosSinVotar{}
	}
	return sist.votosFinales, sist.tamanoVotos, errFinal
}

// Busca el padron en toda la lista de padrones ingresados con busqueda binaria.
func dniEnPadron(dni int, sist *sistemaDeVotacionImplementacion) (bool, int) {
	return _dniEnPadron(sist.padronesRegistrados, dni, 0, len(sist.padronesRegistrados)-1)
}

func _dniEnPadron(arr []int, dni int, init int, fin int) (bool, int) {
	if init > fin {
		return false, -1
	}
	mid := (init + fin) / 2
	if arr[mid] == dni {
		return true, mid

	} else if arr[mid] < dni {
		return _dniEnPadron(arr, dni, mid+1, fin)
	} else {
		return _dniEnPadron(arr, dni, init, mid-1)
	}
}
