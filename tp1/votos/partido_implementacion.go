package votos

import (
	"strconv"
)

type partidoImplementacion struct {
	nombre     string
	candidatos [CANT_VOTACION]string
	votos      [3]int
}

type partidoEnBlanco struct {
	nombre string
	votos  [3]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.candidatos = candidatos
	return partido
}

func CrearVotosEnBlanco() Partido {
	partido := new(partidoEnBlanco)
	partido.nombre = "Votos en Blanco"
	return partido
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	cadenaAux := " votos"
	if partido.votos[tipo] == 1 {
		cadenaAux = " voto"
	}
	return partido.nombre + " - " + partido.candidatos[tipo] + ": " + strconv.Itoa(partido.votos[tipo]) + cadenaAux
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	cadenaAux := " votos"
	if blanco.votos[tipo] == 1 {
		cadenaAux = " voto"
	}
	return blanco.nombre + ": " + strconv.Itoa(blanco.votos[tipo]) + cadenaAux
}
