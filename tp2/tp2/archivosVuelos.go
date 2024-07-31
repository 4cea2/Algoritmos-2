package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	Error "tp2/errores"
	SistVuelos "tp2/sistemavuelos"
)

// agregar_archivo <nombre_archivo>: procesa de forma completa un archivo de .csv que contiene datos de vuelos.
func SistemaAgregarArchivo(sist SistVuelos.SistemaDeVuelos, comando string) {
	vuelosCSV, sePudoLeer := AperturaArchivos(comando)
	if !sePudoLeer {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "agregar_archivo"}.Error())
		return
	}
	arregloVuelos, tamanoVuelos := LecturaArchivocsv(vuelosCSV)
	vuelos := asignacionVuelos(arregloVuelos, tamanoVuelos)
	sist.ComandoCargarVuelos(vuelos[:tamanoVuelos])
	fmt.Fprintf(os.Stdout, "OK\n")
}

// ver_tablero <K cantidad vuelos> <modo: asc/desc> <desde> <hasta>: muestra los K vuelos ordenados por fecha de
// forma ascendente (asc) o descendente (desc), cuya fecha de despegue esté dentro de el intervalo <desde> <hasta> (inclusive).
func SistemaVerTablero(sist SistVuelos.SistemaDeVuelos, comandos []string) {
	if len(comandos) != 4 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "ver_tablero"}.Error())
		return
	}
	k := comandos[0]
	modo := comandos[1]
	desde := comandos[2]
	hasta := comandos[3]
	cantVuelos, _ := strconv.Atoi(k)

	if (cantVuelos < 0) || strings.Compare(desde, hasta) > 0 {
		//fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{"ver_tablero"}.Error())
		fmt.Fprintf(os.Stdout, "OK\n")
		return
	}
	vuelosTabla, errTabla := sist.ComandoVerTablero(cantVuelos, modo, desde, hasta)
	if errTabla != nil {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "ver_tablero"}.Error())
		return
	}
	for i := 0; i < len(vuelosTabla); i++ {
		vueloVisualizado := fmt.Sprint(vuelosTabla[i].FechaPartida, " - ", vuelosTabla[i].NumeroVuelo)
		fmt.Fprintf(os.Stdout, "%s\n", vueloVisualizado)
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

// info_vuelo <código vuelo>: muestra toda la información posible en sobre el vuelo que tiene el código pasado por parámetro
func SistemaInfoVuelo(sist SistVuelos.SistemaDeVuelos, comandos []string) {
	if len(comandos) != 1 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "info_vuelo"}.Error())
		return
	}
	nro := comandos[0]
	numeroVuelo, err := strconv.Atoi(nro)
	if err != nil || numeroVuelo < 0 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "info_vuelo"}.Error())
		return
	}
	vueloFinal, errInfoVuelo := sist.ComandoInfoVuelo(numeroVuelo)
	if errInfoVuelo != nil {
		fmt.Fprintf(os.Stderr, "%s\n", errInfoVuelo)
		return
	}
	mostrarVuelo(vueloFinal)
	fmt.Fprintf(os.Stdout, "OK\n")
}

// prioridad_vuelos <K cantidad vuelos>: muestra los códigos de los K vuelos que tienen mayor prioridad
func SistemaPrioridadVuelos(sist SistVuelos.SistemaDeVuelos, k string) {
	cantidadVuelos, err := strconv.Atoi(k)
	if err != nil || cantidadVuelos < 0 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "prioridad_vuelos"}.Error())
		return
	}

	vuelosPrioritarios := sist.ComandoPrioridadVuelos(cantidadVuelos)

	for i := 0; i < len(vuelosPrioritarios); i++ {
		vueloVisualizado := fmt.Sprint(vuelosPrioritarios[i].Prioridad, " - ", vuelosPrioritarios[i].NumeroVuelo)
		fmt.Fprintf(os.Stdout, "%s\n", vueloVisualizado)
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

// siguiente_vuelo <aeropuerto origen> <aeropuerto destino> <fecha>: muestra la información del vuelo
func SistemaSiguienteVuelo(sist SistVuelos.SistemaDeVuelos, comandos []string) {
	if len(comandos) != 3 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "siguiente_vuelo"}.Error())
		return
	}
	origen := comandos[0]
	destino := comandos[1]
	fecha := comandos[2]
	vueloSiguiente, errSig := sist.ComandoSiguienteVuelos(origen, destino, fecha)
	if errSig != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errSig)

	} else {
		mostrarVuelo(vueloSiguiente)
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

// borrar <desde> <hasta>: borra todos los vuelos cuya fecha de despegue estén dentro del intervalo <desde> <hasta> (inclusive).
func SistemaBorrar(sist SistVuelos.SistemaDeVuelos, comandos []string) {
	if len(comandos) != 2 {
		fmt.Fprintf(os.Stderr, "%s\n", Error.ErrorEnComando{Comando: "borrar"}.Error())
		return
	}
	desde := comandos[0]
	hasta := comandos[1]
	if strings.Compare(desde, hasta) > 0 {
		//fmt.Fprintf(os.Stderr, "%s\n", errBorrar)
		fmt.Fprintf(os.Stdout, "OK\n")
		return
	}
	vuelosBorrados := sist.ComandosBorrar(desde, hasta)
	for i := 0; i < len(vuelosBorrados); i++ {
		mostrarVuelo(vuelosBorrados[i])
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

//+------------------------------------------------------------------------------------------------------------------------------+
//										  #__FUNCIONES_AUXILIARES__#															 +
//+------------------------------------------------------------------------------------------------------------------------------+

func asignacionVuelos(arreglo [][]string, tamanoArreglo int) []SistVuelos.Vuelo {
	vuelos := make([]SistVuelos.Vuelo, tamanoArreglo)
	for i := 0; i < tamanoArreglo; i++ {
		vuelos[i].NumeroVuelo, _ = strconv.Atoi(arreglo[i][0])
		vuelos[i].Aerolinea = arreglo[i][1]
		vuelos[i].Origen = arreglo[i][2]
		vuelos[i].Destino = arreglo[i][3]
		vuelos[i].NumeroCola = arreglo[i][4]
		vuelos[i].Prioridad, _ = strconv.Atoi(arreglo[i][5])
		vuelos[i].FechaPartida = arreglo[i][6]
		vuelos[i].RetrasoSalida, _ = strconv.Atoi(arreglo[i][7])
		vuelos[i].TiempoVuelo = arreglo[i][8]
		vuelos[i].Cancelado = arreglo[i][9]
	}
	return vuelos
}

func mostrarVuelo(vuelo SistVuelos.Vuelo) {
	fmt.Fprintf(os.Stdout, "%d %s %s %s %s %d %s %d %s %s\n", vuelo.NumeroVuelo, vuelo.Aerolinea, vuelo.Origen, vuelo.Destino, vuelo.NumeroCola, vuelo.Prioridad, vuelo.FechaPartida, vuelo.RetrasoSalida, vuelo.TiempoVuelo, vuelo.Cancelado)
}
