package votos_test

import (
	TDAvotante "rerepolez/votos"
	"testing"

	"github.com/stretchr/testify/require"
)

const DNI int = 12345678

func TestVotanteCreacion(t *testing.T) {
	t.Log("Test de creacion de un votante")
	votante := TDAvotante.CrearVotante(DNI)
	require.EqualValues(t, DNI, votante.LeerDNI())
}

func TestVotanteVotaNormal(t *testing.T) {
	t.Log("Test de votacion normal")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	require.EqualValues(t, nil, votante.Votar(TDAvotante.PRESIDENTE, 3))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.GOBERNADOR, 1))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 2))
	votoFinal, err := votante.FinVoto()
	require.Equal(t, nil, err)
	require.Equal(t, 3, votoFinal.VotoPorTipo[TDAvotante.PRESIDENTE])
	require.Equal(t, 1, votoFinal.VotoPorTipo[TDAvotante.GOBERNADOR])
	require.Equal(t, 2, votoFinal.VotoPorTipo[TDAvotante.INTENDENTE])
}

// Se podria llamar tambien TestVotanteFinalizaElVotoYSinHaberVotado, ya que si no votas y lo finalizas, quedas en blanco.
func TestVotanteVotaEnBlanco(t *testing.T) {
	t.Log("Test de votacion en blanco")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	votoFinal, err := votante.FinVoto()
	require.Equal(t, nil, err)
	require.Equal(t, 0, votoFinal.VotoPorTipo[TDAvotante.PRESIDENTE])
	require.Equal(t, 0, votoFinal.VotoPorTipo[TDAvotante.GOBERNADOR])
	require.Equal(t, 0, votoFinal.VotoPorTipo[TDAvotante.INTENDENTE])
}

func TestVotanteVotaLuegoDeshaceUnaVez(t *testing.T) {
	t.Log("Test de deshacer la ultima votacion ya hecha")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	require.EqualValues(t, nil, votante.Votar(TDAvotante.PRESIDENTE, 3))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.GOBERNADOR, 1))
	require.EqualValues(t, nil, votante.Deshacer())
	votoFinal, err := votante.FinVoto()
	require.Equal(t, nil, err)
	require.Equal(t, 3, votoFinal.VotoPorTipo[TDAvotante.PRESIDENTE])
	require.Equal(t, 0, votoFinal.VotoPorTipo[TDAvotante.GOBERNADOR])
}

func TestVotanteVotaLuegoDeshaceTodo(t *testing.T) {
	t.Log("Test de deshacer toda la votacion ya hecha")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	require.EqualValues(t, nil, votante.Votar(TDAvotante.PRESIDENTE, 3))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.GOBERNADOR, 1))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.GOBERNADOR, 5))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.PRESIDENTE, 6))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 2))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.PRESIDENTE, 7))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 9))
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	votoFinal, err := votante.FinVoto()
	require.Equal(t, nil, err)
	require.Equal(t, 0, votoFinal.VotoPorTipo[TDAvotante.PRESIDENTE])
	require.Equal(t, 0, votoFinal.VotoPorTipo[TDAvotante.GOBERNADOR])
	require.Equal(t, 0, votoFinal.VotoPorTipo[TDAvotante.INTENDENTE])
}

func TestVotanteNoVotaYDeshace(t *testing.T) {
	t.Log("Test de deshacer una votacion sin haber votado")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	err := votante.Deshacer()
	require.EqualError(t, err, "ERROR: Sin voto a deshacer")
}

func TestVotanteFinalizaElVotoYDeshace(t *testing.T) {
	t.Log("Test de deshacer una votacion ya finalizada")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 9))
	votoFinal, errFinVoto := votante.FinVoto()
	require.Equal(t, nil, errFinVoto)
	require.Equal(t, 9, votoFinal.VotoPorTipo[TDAvotante.INTENDENTE])
	errDeshacer := votante.Deshacer()
	require.EqualError(t, errDeshacer, "ERROR: Votante FRAUDULENTO: 12345678")
}

func TestVotanteFinalizaElVotoYVotaDeNuevo(t *testing.T) {
	t.Log("Test de votar luego de haber finalizado su voto anteriormente, se lo tiene que descartar")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 9))
	votoFinal, errFinVoto := votante.FinVoto()
	require.Equal(t, nil, errFinVoto)
	require.Equal(t, 9, votoFinal.VotoPorTipo[TDAvotante.INTENDENTE])
	errVotarDeNuevo := votante.Votar(TDAvotante.INTENDENTE, 5)
	require.EqualError(t, errVotarDeNuevo, "ERROR: Votante FRAUDULENTO: 12345678")
	votoFinal, errFinVoto = votante.FinVoto()
	require.EqualError(t, errFinVoto, "ERROR: Votante FRAUDULENTO: 12345678")
	require.False(t, votoFinal.Impugnado)
}

func TestVotanteDeshacerVotosImpugnado(t *testing.T) {
	t.Log("Test de deshacer los votos impugnados hasta llegar al voto no impugnado")
	votante := TDAvotante.CrearVotante(DNI)
	TestVotanteCreacion(t)
	require.EqualValues(t, nil, votante.Votar(TDAvotante.PRESIDENTE, 5))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.GOBERNADOR, 2))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 7))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 0))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.GOBERNADOR, 0))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 0))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.INTENDENTE, 0))
	require.EqualValues(t, nil, votante.Votar(TDAvotante.PRESIDENTE, 0))
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	require.EqualValues(t, nil, votante.Deshacer())
	votoFinal, err := votante.FinVoto()
	require.False(t, votoFinal.Impugnado)
	require.Equal(t, nil, err)
}

/*
Presidente:
Votos en Blanco: 1 voto
Frente para la Derrota - Alan Informaión: 1 voto
Frente para la Derrota - Ignacio Líneas: 0 votos
Frente para la Derrota - Esteban Federico: 1 voto
Pre - Jesús Sintasi: 0 votos
Pre - Javier Colsión: 0 votos
Pre - Yennifer Woolite: 0 votos
++A - Ignacio Gundersen: 1 voto
++A - Emisor Bello: 2 votos
++A - Daniela Peligro: 1 voto
*/

func TestPartidoCreacion(t *testing.T) {
	t.Log("Test de creacion de un partido")
	partido := TDAvotante.CrearPartido("Frente para la Derrota", [TDAvotante.CANT_VOTACION]string{"Alan Informaión", "Ignacio Líneas", "Esteban Federico"})
	require.Equal(t, "Frente para la Derrota - Alan Informaión: 0 votos", partido.ObtenerResultado(TDAvotante.PRESIDENTE))
	require.Equal(t, "Frente para la Derrota - Ignacio Líneas: 0 votos", partido.ObtenerResultado(TDAvotante.GOBERNADOR))
	require.Equal(t, "Frente para la Derrota - Esteban Federico: 0 votos", partido.ObtenerResultado(TDAvotante.INTENDENTE))
}

func TestVotosEnBlancosCreacion(t *testing.T) {
	t.Log("Test de creacion de los votos en blanco")
	partido := TDAvotante.CrearVotosEnBlanco()
	require.Equal(t, "Votos en Blanco: 0 votos", partido.ObtenerResultado(TDAvotante.PRESIDENTE))
	require.Equal(t, "Votos en Blanco: 0 votos", partido.ObtenerResultado(TDAvotante.GOBERNADOR))
	require.Equal(t, "Votos en Blanco: 0 votos", partido.ObtenerResultado(TDAvotante.INTENDENTE))
}

func TestPartidoVotado(t *testing.T) {
	t.Log("Test de votacion a un partido")
	partido := TDAvotante.CrearPartido("Frente para la Derrota", [TDAvotante.CANT_VOTACION]string{"Alan Informaión", "Ignacio Líneas", "Esteban Federico"})
	partido.VotadoPara(TDAvotante.PRESIDENTE)
	partido.VotadoPara(TDAvotante.PRESIDENTE)
	partido.VotadoPara(TDAvotante.PRESIDENTE)
	partido.VotadoPara(TDAvotante.GOBERNADOR)
	partido.VotadoPara(TDAvotante.INTENDENTE)
	partido.VotadoPara(TDAvotante.INTENDENTE)
	require.Equal(t, "Frente para la Derrota - Alan Informaión: 3 votos", partido.ObtenerResultado(TDAvotante.PRESIDENTE))
	require.Equal(t, "Frente para la Derrota - Ignacio Líneas: 1 voto", partido.ObtenerResultado(TDAvotante.GOBERNADOR))
	require.Equal(t, "Frente para la Derrota - Esteban Federico: 2 votos", partido.ObtenerResultado(TDAvotante.INTENDENTE))
}

func TestVotosEnBlancoVotados(t *testing.T) {
	t.Log("Test de votacion en blanco")
	partido := TDAvotante.CrearVotosEnBlanco()
	partido.VotadoPara(TDAvotante.PRESIDENTE)
	partido.VotadoPara(TDAvotante.GOBERNADOR)
	partido.VotadoPara(TDAvotante.GOBERNADOR)
	partido.VotadoPara(TDAvotante.INTENDENTE)
	partido.VotadoPara(TDAvotante.INTENDENTE)
	partido.VotadoPara(TDAvotante.INTENDENTE)
	require.Equal(t, "Votos en Blanco: 1 voto", partido.ObtenerResultado(TDAvotante.PRESIDENTE))
	require.Equal(t, "Votos en Blanco: 2 votos", partido.ObtenerResultado(TDAvotante.GOBERNADOR))
	require.Equal(t, "Votos en Blanco: 3 votos", partido.ObtenerResultado(TDAvotante.INTENDENTE))
}
