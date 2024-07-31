package diccionario_test

import (
	"fmt"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN2 = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que un diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](func(k1 string, k2 string) int { return int(k1[0]) - int(k2[0]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un diccionario vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](func(k1 string, k2 string) int { return int(k1[0]) - int(k2[0]) })
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })
	dicNum := TDADiccionario.CrearABB[int, string](func(k1 int, k2 int) int { return (k1) - (k2) })
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElement2(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](func(k1 string, k2 string) int { return int(k1[0]) - int(k2[0]) })
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioOrdenadoGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDato2(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDatoHopscotch2(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](func(k1 int, k2 int) int { return (k1) - (k2) })
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioBorrar2(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorrados2(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericas2(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](func(k1 int, k2 int) int { return (k1) - (k2) })
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClaveVacia2(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo2(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestCadenaLargaParticular2(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func buscar2(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves2(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar2(cs[0], claves))
	require.NotEqualValues(t, -1, buscar2(cs[1], claves))
	require.NotEqualValues(t, -1, buscar2(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores2(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoValoresConBorrados2(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIterarInternoConCondicionDeCorte(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](func(k1 string, k2 string) int { return int(k1[0]) - int(k2[0]) })
	// Insertar elementos en el diccionario
	abb.Guardar("a", "rosita")
	abb.Guardar("b", "pasame")
	abb.Guardar("c", "la")
	abb.Guardar("d", "prueba")
	abb.Guardar("e", "porfavaaaaaaar")
	/*
			    a
		         \
		          b
		         / \
		        c   d
		             \
		              e

	*/

	visitas := make([]string, 0)
	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false
	// Definir una función para visitar los elementos y realizar la prueba
	// Realizar el recorrido del diccionario utilizando la función Iterar
	abb.Iterar(func(clave string, dato string) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if clave == "d" {
			seguirEjecutando = false
			return false
		}
		visitas = append(visitas, clave)
		return true
	})
	require.Equal(t, "abc", strings.Join(visitas, ""))
	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")

	//Ejemplo de la practica
	abb2 := TDADiccionario.CrearABB[int, string](func(k1 int, k2 int) int { return (k1) - (k2) })
	abb2.Guardar(6, "seis")
	abb2.Guardar(15, "quince")
	abb2.Guardar(10, "diez")
	abb2.Guardar(1, "uno")
	abb2.Guardar(13, "trece")
	abb2.Guardar(11, "once")
	abb2.Guardar(14, "catorce")
	abb2.Guardar(8, "ocho")
	abb2.Guardar(16, "dieciseis")
	abb2.Guardar(4, "cuatro")
	//clavesInOrder := [10]int{1, 4, 6, 8, 10, 11, 13, 14, 15, 16}
	suma := 0
	suma_ptr := &suma
	abb2.Iterar(func(clave int, dato string) bool {
		if clave > 14 {
			return false
		}
		*suma_ptr = *suma_ptr + clave
		return true
	})
	require.Equal(t, 1+4+6+8+10+11+13+14, *suma_ptr)
}

func TestIterarDiccionarioVacio2(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar2(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()
	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar2(primero, claves))
	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar2(segundo, claves))
	require.EqualValues(t, valores[buscar2(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar2(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal2(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar2(primero, claves))
	require.NotEqualValues(t, -1, buscar2(segundo, claves))
	require.NotEqualValues(t, -1, buscar2(tercero, claves))
}

func TestPruebaIterarTrasBorrados2(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestIterarRango(t *testing.T) {
	t.Log("Teste de iterar con rango, hacemos varias pruebas con diferentes valores de 'desde' y 'hasta'")
	abb := TDADiccionario.CrearABB[int, string](func(k1 int, k2 int) int { return (k1) - (k2) })
	abb.Guardar(3, "Tres")
	abb.Guardar(1, "Uno")
	abb.Guardar(4, "Cuatro")
	abb.Guardar(2, "Dos")
	abb.Guardar(5, "Cinco")

	visitas := make([]string, 0)
	desde, hasta := 1, 4
	abb.IterarRango(&desde, &hasta, func(clave int, dato string) bool {
		visitas = append(visitas, fmt.Sprintf("%s:%s", strconv.Itoa(clave), dato))
		return true
	})

	require.Equal(t, []string{"1:Uno", "2:Dos", "3:Tres", "4:Cuatro"}, visitas)

	visitas = make([]string, 0)
	desde, hasta = 2, 5
	abb.IterarRango(&desde, &hasta, func(clave int, dato string) bool {
		visitas = append(visitas, fmt.Sprintf("%s:%s", strconv.Itoa(clave), dato))
		return true
	})

	require.Equal(t, []string{"2:Dos", "3:Tres", "4:Cuatro", "5:Cinco"}, visitas)

	visitas = make([]string, 0)
	hasta = 3
	abb.IterarRango(nil, &hasta, func(clave int, dato string) bool {
		visitas = append(visitas, fmt.Sprintf("%s:%s", strconv.Itoa(clave), dato))
		return true
	})

	require.Equal(t, []string{"1:Uno", "2:Dos", "3:Tres"}, visitas)

	visitas = make([]string, 0)
	desde = 2
	abb.IterarRango(&desde, nil, func(clave int, dato string) bool {
		visitas = append(visitas, fmt.Sprintf("%s:%s", strconv.Itoa(clave), dato))
		return true
	})

	require.Equal(t, []string{"2:Dos", "3:Tres", "4:Cuatro", "5:Cinco"}, visitas)

	visitas = make([]string, 0)
	abb.IterarRango(nil, nil, func(clave int, dato string) bool {
		visitas = append(visitas, fmt.Sprintf("%s:%s", strconv.Itoa(clave), dato))
		return true
	})

	require.Equal(t, []string{"1:Uno", "2:Dos", "3:Tres", "4:Cuatro", "5:Cinco"}, visitas)
}

func TestIteradorRango(t *testing.T) {
	t.Log("Test del iterador con rango de diferentes valores")
	dic := TDADiccionario.CrearABB[int, string](func(k1 int, k2 int) int { return (k1) - (k2) })
	dic.Guardar(6, "seis")
	dic.Guardar(15, "quince")
	dic.Guardar(10, "diez")
	dic.Guardar(1, "uno")
	dic.Guardar(13, "trece")
	dic.Guardar(11, "once")
	dic.Guardar(14, "catorce")
	dic.Guardar(8, "ocho")
	dic.Guardar(16, "dieciseis")
	dic.Guardar(4, "cuatro")
	clavesInOrder := [10]int{1, 4, 6, 8, 10, 11, 13, 14, 15, 16}

	desde, hasta := 1, 16
	iterRango := dic.IteradorRango(&desde, &hasta)
	for i := 0; i < 10; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	desde, hasta = 1, 6
	iterRango = dic.IteradorRango(&desde, &hasta)
	for i := 0; i < 3; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	desde, hasta = 4, 6
	iterRango = dic.IteradorRango(&desde, &hasta)
	for i := 1; i < 3; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	desde, hasta = 6, 8
	iterRango = dic.IteradorRango(&desde, &hasta)
	for i := 2; i < 4; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	desde, hasta = 8, 10
	iterRango = dic.IteradorRango(&desde, &hasta)
	for i := 3; i < 5; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	desde, hasta = 8, 16
	iterRango = dic.IteradorRango(&desde, &hasta)
	for i := 3; i < 10; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	iterRango = dic.IteradorRango(nil, nil)
	for i := 0; i < 10; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	hasta = 13
	iterRango = dic.IteradorRango(nil, &hasta)
	for i := 0; i < 7; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })

	desde = 8
	iterRango = dic.IteradorRango(&desde, nil)
	for i := 3; i < 10; i++ {
		require.True(t, iterRango.HaySiguiente())
		claveActual, _ := iterRango.VerActual()
		require.Equal(t, clavesInOrder[i], claveActual)
		iterRango.Siguiente()
	}
	require.False(t, iterRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterRango.Siguiente() })
}
