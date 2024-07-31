package filaVotos_test

import (
	Error "rerepolez/errores"
	TDAsistema "rerepolez/filaVotos"
	votante "rerepolez/votos"
	"testing"

	"github.com/stretchr/testify/require"
)

//const DNI int = 32000001 //12345678

func TestCrearSistemaVotacion(t *testing.T) {
	t.Log("Test de crear sistema de votacion")
	padrones := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	var votoN2 votante.Voto
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosVacios := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosVacios, votosFinales)
	require.Equal(t, nil, errFinal)
	require.Equal(t, 0, tamanoVotos)

}

func TestSistemaIngresarDniRegistrado(t *testing.T) {
	t.Log("Test de ingresar un dni cargado al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
}
func TestSistemaIngresarDniNoRegistrado(t *testing.T) {
	t.Log("Test de ingresar un dni no cargado al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.EqualError(t, sistema.ComandoIngresar(0), Error.DNIFueraPadron{}.Error())
}
func TestSistemaIngresarDniYaIngresado(t *testing.T) {
	t.Log("Test de ingresar un dni ya ingresado al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
}

func TestSistemaVotacionVotaUnoYFinaliza(t *testing.T) {
	t.Log("Test de registar 1 voto al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 3))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	votoN1.VotoPorTipo[votante.PRESIDENTE] = 1
	votoN1.VotoPorTipo[votante.GOBERNADOR] = 2
	votoN1.VotoPorTipo[votante.INTENDENTE] = 3
	var votoN2 votante.Voto
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 1, tamanoVotos)
	require.Equal(t, nil, errFinal)
}

func TestSistemaVotacionVotaUnoYNoFinaliza(t *testing.T) {
	t.Log("Test de registrar 0 votos aunque hayan votado al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 3))
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	var votoN2 votante.Voto
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 0, tamanoVotos)
	require.EqualError(t, errFinal, Error.ErrorCiudadanosSinVotar{}.Error())
}

func TestSistemaVotacionVotaUnoYDeshaceYFinaliza(t *testing.T) {
	t.Log("Test de registrar 1 voto (modificado por deshacer) al sistema, quedaria en blanco")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 3))
	require.Equal(t, nil, sistema.ComandoDeshacer())
	require.Equal(t, nil, sistema.ComandoDeshacer())
	require.Equal(t, nil, sistema.ComandoDeshacer())
	require.Equal(t, nil, sistema.ComandoFinVotar())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	var votoN2 votante.Voto
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 1, tamanoVotos)
	require.Equal(t, nil, errFinal)
}

func TestSistemaVotacionVotaUnoYDeshaceYNoFinaliza(t *testing.T) {
	t.Log("Test de registrar 0 votos (modificado por deshacer) al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 3))
	require.Equal(t, nil, sistema.ComandoDeshacer())
	require.Equal(t, nil, sistema.ComandoDeshacer())
	require.Equal(t, nil, sistema.ComandoDeshacer())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	var votoN2 votante.Voto
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 0, tamanoVotos)
	require.EqualError(t, errFinal, Error.ErrorCiudadanosSinVotar{}.Error())
}

func TestSistemaVotacionVotoImpugnado(t *testing.T) {
	t.Log("Test de registrar 2 votos impugnado al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 0))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 0))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	votoN1.Impugnado = true
	var votoN2 votante.Voto
	votoN2.Impugnado = true
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 2, tamanoVotos)
	require.Equal(t, nil, errFinal)
}

func TestSistemaVotacionVotoImpugnadoYDeshace(t *testing.T) {
	t.Log("Test de registrar 3 votos desimpugnado (modificado con deshacer) al sistema")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 3))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 1))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 0))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 0))
	require.Equal(t, nil, sistema.ComandoDeshacer())
	require.Equal(t, nil, sistema.ComandoDeshacer())
	require.Equal(t, nil, sistema.ComandoFinVotar())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	votoN1.VotoPorTipo[votante.PRESIDENTE] = 3
	votoN1.VotoPorTipo[votante.GOBERNADOR] = 2
	votoN1.VotoPorTipo[votante.INTENDENTE] = 1
	var votoN2 votante.Voto
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 1, tamanoVotos)
	require.False(t, votosEsperados[0].Impugnado)
	require.Equal(t, nil, errFinal)
}

func TestSistemaVotacionVotoFraudulento(t *testing.T) {
	t.Log("Test de votacion de un fraudulento, se lo descarta de la fila (basicamente lo patea)")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(2))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 3))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 1))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 4))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 5))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 6))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	require.Error(t, sistema.ComandoVotar(votante.INTENDENTE, 10), Error.ErrorVotanteFraudulento{Dni: 1}.Error())
	require.Error(t, sistema.ComandoVotar(votante.INTENDENTE, 10), Error.ErrorVotanteFraudulento{Dni: 2}.Error())
	require.Error(t, sistema.ComandoVotar(votante.INTENDENTE, 10), Error.FilaVacia{}.Error())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	votoN1.VotoPorTipo[votante.PRESIDENTE] = 3
	votoN1.VotoPorTipo[votante.GOBERNADOR] = 2
	votoN1.VotoPorTipo[votante.INTENDENTE] = 1
	var votoN2 votante.Voto
	votoN2.VotoPorTipo[votante.PRESIDENTE] = 4
	votoN2.VotoPorTipo[votante.GOBERNADOR] = 5
	votoN2.VotoPorTipo[votante.INTENDENTE] = 6
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 2, tamanoVotos)
	require.Equal(t, nil, errFinal)
}

func TestSistemaVotacionDeshaceFraudulento(t *testing.T) {
	t.Log("Test de deshacer de un fraudulento, se lo descarta de la fila (basicamente lo patea)")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(2))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 3))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 1))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 4))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 5))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 6))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	require.Error(t, sistema.ComandoDeshacer(), Error.ErrorVotanteFraudulento{Dni: 1}.Error())
	require.Error(t, sistema.ComandoDeshacer(), Error.ErrorVotanteFraudulento{Dni: 2}.Error())
	require.Error(t, sistema.ComandoDeshacer(), Error.FilaVacia{}.Error())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	votoN1.VotoPorTipo[votante.PRESIDENTE] = 3
	votoN1.VotoPorTipo[votante.GOBERNADOR] = 2
	votoN1.VotoPorTipo[votante.INTENDENTE] = 1
	var votoN2 votante.Voto
	votoN2.VotoPorTipo[votante.PRESIDENTE] = 4
	votoN2.VotoPorTipo[votante.GOBERNADOR] = 5
	votoN2.VotoPorTipo[votante.INTENDENTE] = 6
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 2, tamanoVotos)
	require.Equal(t, nil, errFinal)
}

func TestSistemaVotacionFinVotoFraudulento(t *testing.T) {
	t.Log("Test de finalizar el voto de un fraudulento, se lo descarta de la fila (basicamente lo patea)")
	padrones := []int{1, 2, 3}
	sistema := TDAsistema.CrearSistemaVotacion(padrones, len(padrones))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(2))
	require.Equal(t, nil, sistema.ComandoIngresar(1))
	require.Equal(t, nil, sistema.ComandoIngresar(2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 3))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 2))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 1))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	require.Equal(t, nil, sistema.ComandoVotar(votante.PRESIDENTE, 4))
	require.Equal(t, nil, sistema.ComandoVotar(votante.GOBERNADOR, 5))
	require.Equal(t, nil, sistema.ComandoVotar(votante.INTENDENTE, 6))
	require.Equal(t, nil, sistema.ComandoFinVotar())
	require.Error(t, sistema.ComandoFinVotar(), Error.ErrorVotanteFraudulento{Dni: 1}.Error())
	require.Error(t, sistema.ComandoFinVotar(), Error.ErrorVotanteFraudulento{Dni: 2}.Error())
	require.Error(t, sistema.ComandoFinVotar(), Error.FilaVacia{}.Error())
	votosFinales, tamanoVotos, errFinal := sistema.ComandoFinalizarVotacion()
	var votoN1 votante.Voto
	votoN1.VotoPorTipo[votante.PRESIDENTE] = 3
	votoN1.VotoPorTipo[votante.GOBERNADOR] = 2
	votoN1.VotoPorTipo[votante.INTENDENTE] = 1
	var votoN2 votante.Voto
	votoN2.VotoPorTipo[votante.PRESIDENTE] = 4
	votoN2.VotoPorTipo[votante.GOBERNADOR] = 5
	votoN2.VotoPorTipo[votante.INTENDENTE] = 6
	var votoN3 votante.Voto
	var votoN4 votante.Voto
	var votoN5 votante.Voto
	votosEsperados := []votante.Voto{votoN1, votoN2, votoN3, votoN4, votoN5}
	require.Equal(t, votosEsperados, votosFinales)
	require.Equal(t, 2, tamanoVotos)
	require.Equal(t, nil, errFinal)
}
