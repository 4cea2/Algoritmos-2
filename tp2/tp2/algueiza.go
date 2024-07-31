package main

import (
	"bufio"
	"fmt"
	"os"
	SistVuelos "tp2/sistemavuelos"
)

const LINEASINICIALES int = 5

func main() {
	sistemaDeVuelos := SistVuelos.CrearSistemaDeVuelos()
	entradaStdin := bufio.NewScanner(os.Stdin)
	for entradaStdin.Scan() {
		comandoLeidos := LeerEntrada(entradaStdin)

		switch comandoLeidos[0] {
		case "agregar_archivo":
			SistemaAgregarArchivo(sistemaDeVuelos, comandoLeidos[1])

		case "ver_tablero":
			SistemaVerTablero(sistemaDeVuelos, comandoLeidos[1:])

		case "info_vuelo":
			SistemaInfoVuelo(sistemaDeVuelos, comandoLeidos[1:])

		case "prioridad_vuelos":
			SistemaPrioridadVuelos(sistemaDeVuelos, comandoLeidos[1])

		case "siguiente_vuelo":
			SistemaSiguienteVuelo(sistemaDeVuelos, comandoLeidos[1:])

		case "borrar":
			SistemaBorrar(sistemaDeVuelos, comandoLeidos[1:])

		default:
			fmt.Println("le erraste pa")
		}
	}
}
