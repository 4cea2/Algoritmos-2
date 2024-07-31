package sistemavuelos

type Vuelo struct {
	NumeroVuelo   int
	Aerolinea     string
	Origen        string
	Destino       string
	NumeroCola    string
	Prioridad     int
	FechaPartida  string
	RetrasoSalida int
	TiempoVuelo   string
	Cancelado     string
}

type SistemaDeVuelos interface {
	//Carga un nuevo array de vuelos al sistema
	ComandoCargarVuelos(vuelos []Vuelo)

	//Devuelve un vuelo segun el numero de vuelo ingresar o error si el vuelo no pertence al sistema.
	ComandoInfoVuelo(numerovuelo int) (Vuelo, error)

	//Devuelve un array con k vuelos con mayor prioridad
	ComandoPrioridadVuelos(k int) []Vuelo

	//Devuelve un array de k vuelos en el rango desde-hasta segun modo "asc" o "des", ascendente o descendete respectivamente
	//precondicion: "desde" > "hasta", k>=0 y modo :{"asc", "desc"}(modificado en slack)
	ComandoVerTablero(k int, modo string, desde string, hasta string) ([]Vuelo, error)

	//muestra el siguiente vuelo segun el origen y destino considerando la fecha inicial.
	//devuelve error si no hay vuelo registrado desde la fecha indicada.
	ComandoSiguienteVuelos(origen, destino, fecha string) (Vuelo, error)

	//Borra los vuelos en el rago de fechas "dede"-"hasta"
	ComandosBorrar(desde, hasta string) []Vuelo
}
