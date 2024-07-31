package errores

import "fmt"

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type DNIError struct{}

func (e DNIError) Error() string {
	return "ERROR: DNI incorrecto"
}

type DNIFueraPadron struct{}

func (e DNIFueraPadron) Error() string {
	return "ERROR: DNI fuera del padrón"
}

type FilaVacia struct{}

func (e FilaVacia) Error() string {
	return "ERROR: Fila vacía"
}

type ErrorVotanteFraudulento struct {
	Dni int
}

func (e ErrorVotanteFraudulento) Error() string {
	return fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", e.Dni)
}

type ErrorTipoVoto struct{}

func (e ErrorTipoVoto) Error() string {
	return "ERROR: Tipo de voto inválido"
}

type ErrorAlternativaInvalida struct{}

func (e ErrorAlternativaInvalida) Error() string {
	return "ERROR: Alternativa inválida"
}

type ErrorNoHayVotosAnteriores struct{}

func (e ErrorNoHayVotosAnteriores) Error() string {
	return "ERROR: Sin voto a deshacer"
}

type ErrorCiudadanosSinVotar struct{}

func (e ErrorCiudadanosSinVotar) Error() string {
	return "ERROR: Ciudadanos sin terminar de votar"
}
