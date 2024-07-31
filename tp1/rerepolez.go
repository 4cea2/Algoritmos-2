package main

import (
	"bufio"
	"fmt"
	"os"
	Error "rerepolez/errores"
	"rerepolez/filaVotos"
	votos "rerepolez/votos"
	"strconv"
	"strings"
)

func main() {
	var args = os.Args
	partidosCSV, padronesTXT, sePudoIniciarElPrograma := IniciarPrograma(&args)
	if !sePudoIniciarElPrograma {
		return
	}
	partidosPresentados, tamanoPartidos, sistemaDeVotacion := crearSistemaDeVotacion(padronesTXT, partidosCSV)
	entradaStdin := bufio.NewScanner(os.Stdin)
	for entradaStdin.Scan() {
		comandoLeido := leerEntrada(entradaStdin)
		switch comandoLeido[0] {
		case "ingresar":
			sistemaIngresar(comandoLeido, sistemaDeVotacion)
		case "votar":
			sistemaVotar(comandoLeido, sistemaDeVotacion)
		case "deshacer":
			sistemaDeshacer(sistemaDeVotacion)
		case "fin-votar":
			sistemaFinVotar(sistemaDeVotacion)
		default:
		}
	}
	FinalizarPrograma(partidosPresentados, tamanoPartidos, sistemaDeVotacion)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////// MANEJO Y LECTURA DE ARCHIVOS  /////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Verifica si los parametros son validos.
// En caso que sean validos, devuelve true. Caso contrario, falso.
func parametroValidos(args *[]string) bool {
	return len(*args) == 3
}

// Lee la entrada de stdin.
// Devuelve la entrada como un comando facil de manipular.
func leerEntrada(entrada *bufio.Scanner) []string {
	linea := entrada.Text()
	comando := strings.Split(strings.TrimSpace(linea), " ")
	return comando
}

// Abre los archivos de los partidos y padrones.
// Devuelve un puntero a esos archivos y true si se pudieron abrir correctamente.Caso contrarui, false.
func aperturaArchivos(args []string) (*os.File, *os.File, bool) {
	var f_csv *os.File
	var f_txt *os.File
	var err error
	f_csv, err = os.Open(args[1])
	if err != nil {
		defer f_csv.Close()
		return f_csv, f_txt, false
	}
	f_txt, err = os.Open(args[2])
	if err != nil {
		defer f_txt.Close()
		return f_csv, f_txt, false
	}
	return f_csv, f_txt, true
}

const PADRONESINICIALES int = 5

// Lee el archivo de padrones y los guarda en memoria.
// Devuelve dicho arreglo con su respectivo tamano.
func lecturaArchivotxt(txt *os.File) (arregloFinal []int, tamano int) {
	arregloPadrones := make([]int, PADRONESINICIALES)
	capacidadPadrones := PADRONESINICIALES
	scanner := bufio.NewScanner(txt)
	var i int
	for i = 0; scanner.Scan(); i++ {
		if i == capacidadPadrones {
			capacidadPadrones *= 2
			nuevoArregloPadrones := make([]int, capacidadPadrones)
			copy(nuevoArregloPadrones, arregloPadrones)
			arregloPadrones = nuevoArregloPadrones
		}
		padronLeido := scanner.Text()
		arregloPadrones[i], _ = strconv.Atoi(padronLeido)
	}
	return arregloPadrones, i
}

const LINEASINICIALES int = 5

// Lee el archivo csv y lo guarda en memoria.
// Devuelve una matriz (tipo string) de i filas (lineas) y 4 columnas (nombre de partido, candidato a presidente, gobernador y intendente).
func lecturaArchivocsv(csv *os.File) (arregloFinal [][]string, tamano int) {
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
	i++
	return arregloLineas, i
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////	ADMINISTRACION DE PARTIDOS Y VOTOS	////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Asigna y devuelve los partidos que se presentaron.
func asignacionDePartidos(arreglo [][]string, tamanoArreglo int) []votos.Partido {
	partidos := make([]votos.Partido, tamanoArreglo)
	partidos[0] = votos.CrearVotosEnBlanco()
	for i := 1; i < tamanoArreglo; i++ {
		candidatos := [votos.CANT_VOTACION]string{arreglo[i-1][1], arreglo[i-1][2], arreglo[i-1][3]}
		partidos[i] = votos.CrearPartido(arreglo[i-1][0], candidatos)
	}
	return partidos
}

// Asigna los votos para los representantes correspondientes de cada partido.
// Devuelve la cantidad de votos impugnados.
func asignacionDeVotos(partidos []votos.Partido, votosFinales []votos.Voto, tamanoVotos int) int {
	impugnacionesTotales := 0
	for i := 0; i < tamanoVotos; i++ {
		if votosFinales[i].Impugnado {
			impugnacionesTotales++
			continue
		}
		partidos[votosFinales[i].VotoPorTipo[votos.PRESIDENTE]].VotadoPara(votos.PRESIDENTE)
		partidos[votosFinales[i].VotoPorTipo[votos.GOBERNADOR]].VotadoPara(votos.GOBERNADOR)
		partidos[votosFinales[i].VotoPorTipo[votos.INTENDENTE]].VotadoPara(votos.INTENDENTE)
	}
	return impugnacionesTotales
}

// Visualiza los partidos con sus respectivos votos.
func verVotosFinalesPartidos(partidos []votos.Partido, cantPartidos int, votosImpugnados int) {
	postuladosPara := [votos.CANT_VOTACION]string{"Presidente:", "Gobernador:", "Intendente:"}
	var i votos.TipoVoto
	for i = 0; i < votos.CANT_VOTACION; i++ {
		fmt.Println(postuladosPara[i])
		for j := 0; j < cantPartidos; j++ {
			fmt.Println(partidos[j].ObtenerResultado(i))
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
	cadenaAux := "votos"
	if votosImpugnados == 1 {
		cadenaAux = "voto"
	}
	fmt.Fprintf(os.Stdout, "Votos Impugnados: %d %s\n", votosImpugnados, cadenaAux)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////  PROGRAMA  //////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Verifica si se pudo iniciar correctamente el programa.
// En caso de haber no haber un error, devuelve los archivos correctamente y true. Caso contrario, devuelve los archivos como nil y false.
func IniciarPrograma(argumentos *[]string) (*os.File, *os.File, bool) {
	if !parametroValidos(argumentos) {
		var errorParametros Error.ErrorParametros
		fmt.Fprintf(os.Stdout, "%s\n", errorParametros.Error())
		return nil, nil, false
	}
	csv, txt, validacionAperturaArchivos := aperturaArchivos(*argumentos)
	if !validacionAperturaArchivos {
		var errorLeerArchivos Error.ErrorLeerArchivo
		fmt.Fprintf(os.Stdout, "%s\n", errorLeerArchivos.Error())
		return nil, nil, false
	}
	return csv, txt, true
}

// Crea el sistema de votacion a partir de los archivos dados.
// Devuelve el sistema de votacion y los partidos presentados con su respectivo tamaño.
func crearSistemaDeVotacion(txt *os.File, csv *os.File) (partidos []votos.Partido, tamPartidos int, sistema filaVotos.SistemaDeVotacion) {
	arregloPadrones, tamanoPadrones := lecturaArchivotxt(txt)
	padronesPresentados := OrdenarPadronesConRadixSort(arregloPadrones, tamanoPadrones)
	arregloPartidos, tamanoPartidos := lecturaArchivocsv(csv)
	partidosPresentados := asignacionDePartidos(arregloPartidos, tamanoPartidos)
	sistemaDeVotacion := filaVotos.CrearSistemaVotacion(padronesPresentados[:tamanoPadrones], tamanoPadrones, tamanoPartidos)
	return partidosPresentados, tamanoPartidos, sistemaDeVotacion
}

// Ingresa al sistema con el padron ingresado y lo valida.
func sistemaIngresar(comando []string, sistemaDeVotacion filaVotos.SistemaDeVotacion) {
	dniIngresado, errorConversion := strconv.Atoi(comando[1])
	if errorConversion != nil {
		errorConversion = Error.DNIError{}
		fmt.Fprintf(os.Stdout, "%s\n", errorConversion.Error())
		return
	}

	errorIngresar := sistemaDeVotacion.ComandoIngresar(dniIngresado)
	if errorIngresar != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errorIngresar.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

// Emite un voto en el sistema y lo valida
func sistemaVotar(comando []string, sistemaDeVotacion filaVotos.SistemaDeVotacion) {
	if !((comando[1] == "Presidente") || (comando[1] == "Gobernador") || (comando[1] == "Intendente")) {
		var errorTipoVoto Error.ErrorTipoVoto
		fmt.Fprintf(os.Stdout, "%s\n", errorTipoVoto.Error())
		return
	}
	listaIngresada := comando[2]
	var tipoDeVoto votos.TipoVoto
	if comando[1] == "Presidente" {
		tipoDeVoto = votos.PRESIDENTE
	}
	if comando[1] == "Gobernador" {
		tipoDeVoto = votos.GOBERNADOR
	}
	if comando[1] == "Intendente" {
		tipoDeVoto = votos.INTENDENTE
	}
	errorVotar := sistemaDeVotacion.ComandoVotar(tipoDeVoto, listaIngresada)
	if errorVotar != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errorVotar.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

// Deshace un voto en el sistema y lo valida.
func sistemaDeshacer(sistemaDeVotacion filaVotos.SistemaDeVotacion) {
	errDeshacer := sistemaDeVotacion.ComandoDeshacer()
	if errDeshacer != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errDeshacer.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

// Finaliza un voto en el sistema y lo valida.
func sistemaFinVotar(sistemaDeVotacion filaVotos.SistemaDeVotacion) {
	errFinVotar := sistemaDeVotacion.ComandoFinVotar()
	if errFinVotar != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errFinVotar.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "OK\n")
}

// Finaliza el programa. Cuenta los votos emitidos correctamente y los visualiza
func FinalizarPrograma(partidosPresentados []votos.Partido, tamanoPartidos int, sistemaDeVotacion filaVotos.SistemaDeVotacion) bool {
	votosFinales, tamanoVotosFinales, errFinal := sistemaDeVotacion.ComandoFinalizarVotacion()
	if errFinal != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errFinal.Error())
	}
	cantidadVotosImpugnados := asignacionDeVotos(partidosPresentados, votosFinales, tamanoVotosFinales)
	verVotosFinalesPartidos(partidosPresentados, tamanoPartidos, cantidadVotosImpugnados)
	return true
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////ORDENAMIENTO////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Ordena un arr de enteros con el metodo counting sort.
// arr: array de enteros; maxDigito: el maximo valor de digito a ordenar, criterio: el criterio para ordenar el arr
func ordenar(arr []int, tamarr int, maxDigito int, criterio func(int) int) []int {
	//counting
	frecuencias := make([]int, maxDigito) //por que le voy a pasar un digito del 0:9
	acum := make([]int, maxDigito)        //por que hay que sumar las frecuencias para obtener los indices
	final := make([]int, len(arr))

	//cuento frecuencias
	for i := 0; i < tamarr; i++ {
		frecuencias[criterio(arr[i])]++
	}

	//suma acumulada para ver los indices principales
	for i := 1; i < len(acum); i++ {
		acum[i] = acum[i-1] + frecuencias[i-1]
	}

	//ordeno y aumento indices.
	for i := 0; i < tamarr; i++ {
		final[acum[criterio(arr[i])]] = arr[i]
		acum[criterio(arr[i])]++
	}
	return final

}

// Se implementa el algortimo radix sort para ordenar padrones de tipo entero
// se usa la funcion auxiliar counting sort para ordenar de mayor a menor un padron de tipo xx.yyy.zzz o xx.yyy.zz
func OrdenarPadronesConRadixSort(arr []int, tam int) []int {
	//ordenar ultimos digitos (dos o tres)
	arr = ordenar(arr, tam, 999+1, func(elem int) int {
		ultimosDigitos := elem % 1000
		if elem >= 100000000 {
			ultimosDigitos = elem % 10000
		}
		return ultimosDigitos
	})

	//Ordenar tres digitos del medio
	arr = ordenar(arr, tam, 999+1, func(elem int) int { return (elem % 1000000) / 1000 })

	//Ordenar primos dos digitos
	arr = ordenar(arr, tam, 99+1, func(elem int) int { return elem / 1000000 })

	return arr
}
