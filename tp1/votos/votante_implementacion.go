package votos

import (
	"rerepolez/errores"
	"tdas/pila"
)

type votanteImplementacion struct {
	dni    int
	votos  pila.Pila[Voto]
	yaVote bool
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.yaVote = false
	votante.dni = dni
	votante.votos = pila.CrearPilaDinamica[Voto]()
	return votante
}

func (votante *votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	//Caso en donde ya finalizo su voto, hay que descartarlo.
	if votante.yaVote {
		return errores.ErrorVotanteFraudulento{votante.LeerDNI()}
	}
	//Caso general si desea votar.
	var votoAEmitir Voto
	if votante.votos.EstaVacia() {

	} else {
		votoAEmitir = votante.votos.VerTope()
	}
	votoAEmitir.VotoPorTipo[tipo] = alternativa
	if alternativa == 0 {
		votoAEmitir.Impugnado = true
	}
	votante.votos.Apilar(votoAEmitir)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	//Caso en donde no vota y deshace.
	if votante.votos.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	//Caso en donde ya finalizo su voto, hay que descartarlo.
	if votante.yaVote {
		return errores.ErrorVotanteFraudulento{votante.dni}
	}
	//Caso general en donde quiere deshacer.
	votante.votos.Desapilar()
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	//Caso en donde el votante ya finalizo su voto, y quiera finalizarlo de nuevo.
	if votante.yaVote {
		return votante.votos.VerTope(), errores.ErrorVotanteFraudulento{votante.dni}
	}
	//Caso en donde el votante no haya votado y quiera finalizar su voto.
	if votante.votos.EstaVacia() {
		var votoEnBlanco Voto
		votante.votos.Apilar(votoEnBlanco)
	}
	votante.yaVote = true
	return votante.votos.VerTope(), nil
}
