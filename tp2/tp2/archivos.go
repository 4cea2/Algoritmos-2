package main

import (
	"bufio"
	"os"
	"strings"
)

// Abre los archivos de los vuelos
// Devuelve un puntero al archivo y true si se pudo abrir correctamente. Caso contrario, false.
func AperturaArchivos(archivo string) (*os.File, bool) {
	var f_csv *os.File
	var err error
	f_csv, err = os.Open(archivo)
	if err != nil {
		defer f_csv.Close()
		return f_csv, false
	}
	return f_csv, true
}

// Lee el archivo csv y lo guarda en memoria.
// Devuelve una matriz (tipo string) de i filas (lineas) y 10 columnas (caracteristicas del vuelo)
func LecturaArchivocsv(csv *os.File) (arregloFinal [][]string, tamano int) {
	arregloLineas := make([][]string, LINEASINICIALES)
	capacidadLineas := LINEASINICIALES
	scanner := bufio.NewScanner(csv)
	var i int
	for i = 0; scanner.Scan(); i++ {
		if i == capacidadLineas {
			capacidadLineas *= 2
			nuevoArregloLineas := make([][]string, capacidadLineas)
			copy(nuevoArregloLineas, arregloLineas)
			arregloLineas = nuevoArregloLineas

		}
		lineaLeida := scanner.Text()
		arregloLineas[i] = strings.Split(strings.TrimSpace(lineaLeida), ",")
	}
	return arregloLineas, i
}

// Lee la entrada de stdin.
// Devuelve la entrada como un comando facil de manipular.
func LeerEntrada(entrada *bufio.Scanner) []string {
	linea := entrada.Text()
	comando := strings.Split(strings.TrimSpace(linea), " ")
	return comando
}
