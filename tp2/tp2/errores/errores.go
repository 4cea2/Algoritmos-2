package errores

import "fmt"

type ErrorEnComando struct {
	Comando string
}

func (e ErrorEnComando) Error() string {
	return fmt.Sprintf("Error en comando %s", e.Comando)
}

type ErrorSiguiente struct {
	Origen  string
	Destino string
	Fecha   string
}

func (e ErrorSiguiente) Error() string {
	return fmt.Sprintf("No hay vuelo registrado desde %s hacia %s desde %s", e.Origen, e.Destino, e.Fecha)
}
