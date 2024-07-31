package sistemavuelos

import (
	"fmt"
	"strings"
	Heap "tdas/cola_prioridad"
	Dic "tdas/diccionario"
	Error "tp2/errores"
)

type sistemaDeVuelosImplementacion struct {
	vuelosInformacionDesordenada Dic.Diccionario[int, Vuelo]            //Vuelos guardados por numero de vuelo (desordenado, con hash todo es O(1))
	vuelosSiguientes             Dic.DiccionarioOrdenado[string, Vuelo] //Vuelos ordenados por fecha de partida y numero de vuelo(comparacion alfanumerica) (ordenado, con abb es O(log(n)))
}

const CANTIDADINICIAL int = 5

func CrearSistemaDeVuelos() SistemaDeVuelos {
	sist := new(sistemaDeVuelosImplementacion) //nrovuelos-fecha.partida
	sist.vuelosInformacionDesordenada = Dic.CrearHash[int, Vuelo]()
	sist.vuelosSiguientes = Dic.CrearABB[string, Vuelo](cmpFechaPartidaNumeroVuelo)

	return sist
}

func (sistema *sistemaDeVuelosImplementacion) ComandoCargarVuelos(vuelos []Vuelo) {

	for _, vuelo := range vuelos {
		if sistema.vuelosInformacionDesordenada.Pertenece(vuelo.NumeroVuelo) {
			sistema.borrarVueloViejo(vuelo.NumeroVuelo)
		}
		if sistema.vuelosSiguientes.Pertenece(vuelo.FechaPartida) {
			sistema.actualizarVuelo(vuelo)
		} else {
			sistema.guardarVueloNuevo(vuelo)
		}
		sistema.vuelosInformacionDesordenada.Guardar(vuelo.NumeroVuelo, vuelo)
	}
}

func (sistema *sistemaDeVuelosImplementacion) ComandoInfoVuelo(numerovuelo int) (Vuelo, error) {
	if sistema.vuelosInformacionDesordenada.Pertenece(numerovuelo) {
		return sistema.vuelosInformacionDesordenada.Obtener(numerovuelo), nil
	} else {
		var vueloVacio Vuelo
		return vueloVacio, Error.ErrorEnComando{Comando: "info_vuelo"}
	}
}

func (sistema *sistemaDeVuelosImplementacion) ComandoSiguienteVuelos(origen, destino, fecha string) (Vuelo, error) {
	claveFecha := concatenarFechaPartidaNumeroVuelo(fecha, 0000)
	var vueloVacio Vuelo
	for iterAbb := sistema.vuelosSiguientes.IteradorRango(&claveFecha, nil); iterAbb.HaySiguiente(); iterAbb.Siguiente() {
		_, vueloActual := iterAbb.VerActual()
		if cmpFechaHoraAbb(vueloActual.FechaPartida, fecha) == 0 {
			continue
		}
		if vueloActual.Origen == origen && vueloActual.Destino == destino {
			return vueloActual, nil
		}
	}
	return vueloVacio, Error.ErrorSiguiente{Origen: origen, Destino: destino, Fecha: fecha}
}

func (sistema *sistemaDeVuelosImplementacion) ComandosBorrar(desde, hasta string) []Vuelo {
	vuelosBorrados := make([]Vuelo, CANTIDADINICIAL)
	capacidadActual := CANTIDADINICIAL
	var i int
	formatoDesde := concatenarFechaPartidaNumeroVuelo(desde, 0000)
	fomateHasta := concatenarFechaPartidaNumeroVuelo(hasta, 9999)

	for iterAbb := sistema.vuelosSiguientes.IteradorRango(&formatoDesde, &fomateHasta); iterAbb.HaySiguiente(); iterAbb.Siguiente() {
		claveActual, vueloActual := iterAbb.VerActual()
		if i == len(vuelosBorrados) {
			capacidadActual = cap(vuelosBorrados)
			capacidadNueva := 2 * capacidadActual
			array := make([]Vuelo, capacidadNueva)
			copy(array, vuelosBorrados)
			vuelosBorrados = array
		}
		vuelosBorrados[i] = sistema.vuelosInformacionDesordenada.Borrar(vueloActual.NumeroVuelo)
		sistema.vuelosSiguientes.Borrar(claveActual)
		i++
	}

	return vuelosBorrados[:i]
}

func (sistema *sistemaDeVuelosImplementacion) ComandoVerTablero(k int, modo string, desde string, hasta string) ([]Vuelo, error) {
	vuelosCantidad := sistema.vuelosSiguientes.Cantidad()
	vuelosFinalesAsc := make([]Vuelo, vuelosCantidad)
	vuelosFinalesDesc := make([]Vuelo, vuelosCantidad)
	var i int
	formatoDesde := concatenarFechaPartidaNumeroVuelo(desde, 0000)
	fomateHasta := concatenarFechaPartidaNumeroVuelo(hasta, 9999)
	for iterAbb := sistema.vuelosSiguientes.IteradorRango(&formatoDesde, &fomateHasta); iterAbb.HaySiguiente() && i < vuelosCantidad; iterAbb.Siguiente() {
		_, vueloActual := iterAbb.VerActual()
		vuelosFinalesAsc[i] = vueloActual
		vuelosFinalesDesc[vuelosCantidad-i-1] = vueloActual
		i++
	}
	//Corta cuando guarda todos los vuelos que haya en 'desde' y 'hasta', lo denoto 'a'.
	//i es la cantidad de vuelos que guarde con el rango.
	if modo == "asc" {
		if i > k {
			return vuelosFinalesAsc[:k], nil
		} else {
			return vuelosFinalesAsc[:i], nil
		}
	} else {
		if i <= k {
			return vuelosFinalesDesc[(vuelosCantidad - i):], nil
		}
		if i > k {
			return vuelosFinalesDesc[(vuelosCantidad - i):(vuelosCantidad - i + k)], nil
		}
		return vuelosFinalesDesc, nil
	}

}

func (sistema *sistemaDeVuelosImplementacion) ComandoPrioridadVuelos(k int) []Vuelo {
	//Obtengo los n vuelos cargados en el sistema, pero en funcion de su prioridad.
	vuelosCantidad := sistema.vuelosInformacionDesordenada.Cantidad()
	vuelosCargadosPorPrioridad := make([]Vuelo, vuelosCantidad)
	//{[P1, N4], [P1, N1]}
	iter := sistema.vuelosInformacionDesordenada.Iterador()
	for i := 0; iter.HaySiguiente(); i++ {
		_, vuelo := iter.VerActual()
		vuelosCargadosPorPrioridad[i] = vuelo
		iter.Siguiente()
	}
	// ~O(n) por ser un hash y operaciones O(1)

	//Los cargo en un heap con arreglo (esto aplica heapfy, osea, O(n))
	heap := Heap.CrearHeapArr[Vuelo](vuelosCargadosPorPrioridad, cmpVuelos)
	cantidadFinal := k
	if k > vuelosCantidad {
		cantidadFinal = vuelosCantidad
	}
	vuelosFinales := make([]Vuelo, cantidadFinal)
	if cantidadFinal == 0 {
		var vuelosVacios []Vuelo
		return vuelosVacios
	}
	primerVuelo := heap.Desencolar()
	vuelosFinales[0] = sistema.vuelosInformacionDesordenada.Obtener(primerVuelo.NumeroVuelo)
	for i := 1; i < cantidadFinal; i++ {
		vueloFinal := heap.Desencolar()
		vuelosFinales[i] = sistema.vuelosInformacionDesordenada.Obtener(vueloFinal.NumeroVuelo)
	}
	//Finalmente, queda O(n + k*log(n))
	return vuelosFinales[:cantidadFinal]
}

//+---------------------------------------------------------------------------------------------------------------------------+
//+------------------------------------------+Funciones Axuliares+------------------------------------------------------------+
//+---------------------------------------------------------------------------------------------------------------------------+

func (sistema *sistemaDeVuelosImplementacion) borrarVueloViejo(nro int) {
	vueloViejo := sistema.vuelosInformacionDesordenada.Obtener(nro)
	claveVieja := concatenarFechaPartidaNumeroVuelo(vueloViejo.FechaPartida, vueloViejo.NumeroVuelo)
	sistema.vuelosSiguientes.Borrar(claveVieja)
}
func (sistema *sistemaDeVuelosImplementacion) guardarVueloNuevo(vuelo Vuelo) {
	claveNueva := concatenarFechaPartidaNumeroVuelo(vuelo.FechaPartida, vuelo.NumeroVuelo)
	sistema.vuelosSiguientes.Guardar(claveNueva, vuelo)
}
func (sistema *sistemaDeVuelosImplementacion) actualizarVuelo(vuelo Vuelo) {
	sistema.borrarVueloViejo(vuelo.NumeroVuelo)
	sistema.guardarVueloNuevo(vuelo)
}

// Compara fechas+hora de tipo string
func cmpFechaHoraAbb(k1, k2 string) int {
	return strings.Compare(k1, k2)
}

// Compara la cadena "FechaPartida-NumeroVuelo" alfanumericamente.
func cmpFechaPartidaNumeroVuelo(k1, k2 string) int {
	return strings.Compare(k1, k2)
}

// Concatena las dos cadenas mediante "."
func concatenarFechaPartidaNumeroVuelo(fecha string, nro int) string {
	return fmt.Sprint(fecha, "-", fmt.Sprint(nro))
}

func cmpVuelos(v1, v2 Vuelo) int {
	if v1.Prioridad == v2.Prioridad {
		nro1vuelostring := fmt.Sprint(v1.NumeroVuelo)
		nro2vuelostring := fmt.Sprint(v2.NumeroVuelo)
		return (-1) * strings.Compare(nro1vuelostring, nro2vuelostring)
	} else {
		return v1.Prioridad - v2.Prioridad
	}
}
