#!/usr/bin/python3
import sys
import csv
from grafo import Grafo
from heap import Heap
from pila import Pila
from biblioteca import (
    camino_minimo_dijkstra,
    camino_minimo_bfs,
    mst_prim,
    centralidad,
    topologico_grados
)

class FlyCombi:

    def __init__(self):
        self.grafo_tiempo = Grafo(es_dirigido=False) #Grafo de aeropuertos con aristas en funcion del tiempo
        self.grafo_precio = Grafo(es_dirigido=False) #Grafo de aeropuertos con aristas en funcion de precio
        self.grafo_frecuencia = Grafo(es_dirigido=False) #Grafo de aeropuertos con aristas en funcion la frecuencia de vuelos
        self.aeropuertos_de_ciudad = {} #Diccionario de aeropuertos que hay en una ciudad
        self.ciudad_del_aeropuerto = {} #Diccionario de la informacion de la ciudad perteneciente al aeropuerto
        self.informacion_conexion = {}  #Diccionario de la informacion de los vuelos
    
    def cargar_aeropuertos(self, aeropuertos):
        """
        Carga las ciudades y en ellas los aeropuertos que contengan.
        """
        for datos in aeropuertos:
            ciudad = datos[0]
            aeropuerto = datos[1]
            self.ciudad_del_aeropuerto[aeropuerto] = datos

            if ciudad not in self.aeropuertos_de_ciudad:
                self.aeropuertos_de_ciudad[ciudad] = [aeropuerto] 
            else:
                self.aeropuertos_de_ciudad[ciudad].append(aeropuerto)  

            self.grafo_tiempo.agregar_vertice(aeropuerto)
            self.grafo_precio.agregar_vertice(aeropuerto)   
            self.grafo_frecuencia.agregar_vertice(aeropuerto) 


    def cargar_conexiones(self, conexiones):
        """
        Carga las aristas con sus respectivos vertices y peso a los grafos.
        """
        cantidad_de_vuelos_totales = 0
        for datos in conexiones:
            cantidad_de_vuelos_totales += float(datos[4])
            
        for datos in conexiones:
            v = datos[0]
            w = datos[1]
            
            peso_tiempo = datos[2]
            peso_precio =  datos[3]
            peso_cantidad_vuelos = datos[4]
            frecuencia = 100*float(peso_cantidad_vuelos)/cantidad_de_vuelos_totales

            self.grafo_tiempo.agregar_arista(v, w, peso_tiempo)
            self.grafo_precio.agregar_arista(v, w, peso_precio)
            self.grafo_frecuencia.agregar_arista(v, w, 1/frecuencia)
            
            self.informacion_conexion[v+"-"+w] = datos
        

    def encontrar_caminos(self, ciudad_origen, ciudad_destino, _criterio):
        """ 
        Devolvemos los aeropuertos con los cuales vamos de la ciudad origen a 
        la ciudad destino de la forma más rápida o más barata, según corresponda
        """
        caminos_minimos = []

        if _criterio == "rapido":
            mi_grafo = self.grafo_tiempo
        elif _criterio == "barato":
            mi_grafo = self.grafo_precio
        else: 
            raise AssertionError("opcion incorrecta, ni rapido ni barato pe")  
         
        for aeropuerto_destino in self.aeropuertos_de_ciudad[ciudad_destino]:
            for aeropuerto_origen in self.aeropuertos_de_ciudad[ciudad_origen]:
                padres, _ = camino_minimo_dijkstra(mi_grafo, aeropuerto_origen)             
                camino = reconstruir_camino(padres, aeropuerto_origen, aeropuerto_destino)
                caminos_minimos.append(camino)
        if _criterio == "barato":
            return camino_con_menos_costo(caminos_minimos, self.informacion_conexion)
        else:
            return min(caminos_minimos, key=len)
   
    
    def camino_escalas(self, ciudad_origen, ciudad_destino):
        """
        Encuentra el camino que menor escalas tenga, sin importar los costos.
        desde ciudad origen a ciudad destino.
        """
        caminos_minimos = []
        mi_grafo = self.grafo_precio
        for aeropuerto_destino in self.aeropuertos_de_ciudad[ciudad_destino]:
            for aeropuerto_origen in self.aeropuertos_de_ciudad[ciudad_origen]:
                padres, _ = camino_minimo_bfs(mi_grafo, aeropuerto_origen)               
                camino = reconstruir_camino(padres, aeropuerto_origen, aeropuerto_destino)
                caminos_minimos.append(camino)
        
        return min(caminos_minimos, key=len)
    
    def k_aeropuertos_importantes(self, k):
        centros = centralidad(self.grafo_frecuencia)
        return top_k(centros, k)
    
    def nueva_aerolinea(self, archivocsv, vuelos):
        mi_grafo = self.grafo_precio
        arbol = mst_prim(mi_grafo)
        vuelos_a_escribir = []
        for v in arbol.obtener_vertices():
            for w in arbol.adyacentes(v):
                clave1 = v+"-"+w
                clave2 = w+"-"+v
                if clave1 in self.informacion_conexion:
                    vuelos_a_escribir.append(self.informacion_conexion[clave1])
                elif clave2 in self.informacion_conexion:
                    vuelos_a_escribir.append(self.informacion_conexion[clave2])
                else:
                    continue
        with open(archivocsv, 'w', newline='') as file:
            writer = csv.writer(file, delimiter=',')
            writer.writerows(vuelos_a_escribir)

    def exportar_kml(self, archivokml, ultimo_camino):
        datos = []
        for aeropuerto in ultimo_camino:
            datos.append(self.ciudad_del_aeropuerto[aeropuerto])
        escribir_kml(archivokml, datos)

    def itinerario(self, csv):
        lineas = cargar_csv(csv)
        nuevo_grafo = Grafo(es_dirigido=True)
        for ciudad in lineas[0]:
            nuevo_grafo.agregar_vertice(ciudad)
        for orden in lineas[1:]:
            nuevo_grafo.agregar_arista(orden[0], orden[1])
        orden_ciudades = topologico_grados(nuevo_grafo)
        print(", ".join(orden_ciudades))
        indice = 0
        while indice < len(orden_ciudades)-1:
            ciudad_i = orden_ciudades[indice]
            ciudad_j = orden_ciudades[indice+1]
            print(" -> ".join(self.camino_escalas(ciudad_i, ciudad_j)))
            indice = indice + 1
        
     
class Errors():
    def __init__(self, error):
        self.error = error
        
    def CaminosMinimosBaratoRapido(self):
        print("Error al encontrar camino:\n", str(self.error))
        
    def CaminosMinimosEscalas(self):
        print("Error: no se pudo obtener las escalas \n", str(self.error))
        
    def KAeropuertoMasImportante(self):
        print("Error: no se pudo obtener el aropuerto mas importante \n", str(self.error))
    
    def OptimizacionNuevaAerolinea(self): 
        print("Error: en cargar nueva aerolinea \n", str(self.error))       

    def ItinerarioCultural(self):
        print("Error: en generar itinerario \n", str(self.error))
    
    def ExportarKML(self):
        print("Error: al exportar kml \n", str(self.error))
   
           
##+---------------------------------------------------------------------------------------------------------+###
##+-----------------------------------_FUNCIONES AUXILIARES_------------------------------------------------+###
##+---------------------------------------------------------------------------------------------------------+###

def top_k(dic, _k):
    """
    Devuelve los k elementos con mas prioridad (int).
    """
    losKmayores = []
    heap = Heap(False)
    heap.heapify(dic)
    for _ in range(_k):
        losKmayores.append(heap.desencolar())
    return losKmayores
        

def reconstruir_camino(padre, inicio, fin):
    """
    Devuelve el camino reconstruido desde 'inicio' hasta 'fin'.
    """
    vertice = fin
    camino = []
    
    while vertice != inicio:
        camino.append(vertice)
        vertice = padre[vertice]

    camino.append(inicio)
    camino.reverse()
    return camino

def camino_con_menos_costo(caminos, conexion_caminos):
    """
    Devuelve el camino con menor precio posible.
    """
    precios_por_camino = {}
    precios_finales = []
    for camino_actual in caminos:
        precio_camino_actual = 0
        i = 0
        while i < (len(camino_actual)-1):
            conex_v1 = camino_actual[i]+"-"+camino_actual[i+1]
            conex_v2 = camino_actual[i+1]+"-"+camino_actual[i]
            if conex_v1 in conexion_caminos:
                precio_conex = conexion_caminos[conex_v1][3]
            else:
                precio_conex = conexion_caminos[conex_v2][3]
            i = i + 1
            precio_camino_actual = precio_camino_actual + float(precio_conex)
        precios_por_camino[precio_camino_actual] = camino_actual
        precios_finales.append(precio_camino_actual)
    return precios_por_camino[min(precios_finales)]



def iniciar_programa(argumentos):
    """
    Verifica si se pudo iniciar correctamente el programa.
    En caso de haber no haber ningun inconveniente, devuelve los aeropuertos y vuelos leidos correctamente y True. Caso contrario, devuelve los archivos como None y False.
    """
    if len(argumentos) != 3:
        return False, None, None
    
    try:
        aeropuertos_csv = argumentos[1]
        vuelos_csv = argumentos[2]

        lista_aeropuertos = cargar_csv(aeropuertos_csv)
        lista_vuelos = cargar_csv(vuelos_csv)
    except:
        return False, None, None

    return True, lista_aeropuertos, lista_vuelos

def cargar_csv(_archivocsv):
    """
    Abre el archivo, lo lee, guarda en memoria y lo cierra.
    Devuelve una lista de las lineas leidas del csv, separados por columnas (',').
    """
    lista_lineas = []
    with open(_archivocsv, 'r', encoding='utf-8') as csv_abierto:
        csv_lectura = csv.reader(csv_abierto)
        for linea_leida in csv_lectura:
            lista_lineas.append(linea_leida)
    return lista_lineas

def leer_stdin(_stdin):
    """
    Lee la entrada de stdin.
    Devuelve la entrada como un comando facil de manipular.
    """
    parametros = _stdin.split(',')
    aux = parametros[0].split(" ")
    comando_final = []
    comando_final.append(aux[0])
    if aux[0] == "camino_escalas" and len(aux[1:]) == 2:
        comando_final.append(aux[1]+ " " +aux[2])
    else:
        comando_final.append(aux[1])
    for aiuda in parametros[1:]:
        comando_final.append(aiuda)
    return comando_final
    
    
def escribir_kml(archivokml, datos):
    """
    Escribe el archivokml en formato kml con sus respectivos datos.
    """
    with open(archivokml, 'w', encoding = 'utf-8') as file:
        file.write("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
        file.write("<kml xmlns=\"http://earth.google.com/kml/2.1\">\n")
        file.write("    <Document>\n")
        file.write("        <name>"+archivokml[:(len(archivokml)-4)]+"</name>\n")
        file.write("        <description>Un ejemplo introductorio para mostrar la sintaxis KML.</description>\n\n")
        for linea in datos:
            file.write("        <Placemark>\n")
            file.write("            <name>"+linea[1]+"</name>\n")
            file.write("            <description>"+linea[0]+"</description>\n")
            file.write("            <Point>\n")
            file.write("                <coordinates>"+linea[3]+", "+linea[2]+"</coordinates>\n")
            file.write("            </Point>\n")
            file.write("        </Placemark>\n\n")

        i = 0
        while i < len(datos)-1:
            file.write("        <Placemark>\n")
            file.write("            <LineString>\n")
            file.write("                <coordinates>"+datos[i][3]+", "+datos[i][2]+" "+datos[i+1][3]+", "+datos[i+1][2]+"</coordinates>\n")
            file.write("            </LineString>\n")
            file.write("        </Placemark>\n\n")
            i = i + 1
        file.write("    </Document>\n")
        file.write("</kml>\n")


##+---------------------------------------------------------------------------------------------------------+###
##+-------------------------------------------------_MAIN_--------------------------------------------------+###
##+---------------------------------------------------------------------------------------------------------+###

if __name__ == "__main__":
    argumentos = sys.argv
    se_pude_iniciar_el_programa, aeropuertos, vuelos = iniciar_programa(argumentos)
    if not se_pude_iniciar_el_programa:
        raise AssertionError("No se pudo iniciar el programa")  
    
    fly = FlyCombi()
    fly.cargar_aeropuertos(aeropuertos)
    fly.cargar_conexiones(vuelos)  
    
    caminos_mas_rapidos = Pila()
    for line in sys.stdin:
        comandos = leer_stdin(line.rstrip())
        
        if comandos[0] == "camino_mas":
            criterio, origen, destino = comandos[1], comandos[2], comandos[3]
            try:
                camino_final = fly.encontrar_caminos(origen, destino, criterio)
                print(" -> ".join(camino_final))
                caminos_mas_rapidos.apilar(camino_final)
            except Exception as e:
                Errors(e).CaminosMinimosBaratoRapido()

        elif comandos[0] == "camino_escalas":
            try:
                origen = comandos[1]
                destino = comandos[2]
                
                camino_final = fly.camino_escalas(origen, destino)
                #caminos_mas_rapidos.apilar(camino_final)
                print(" -> ".join(camino_final))
            except Exception as e:
                Errors(e).CaminosMinimosEscalas()
            
        elif comandos[0] == "centralidad":
            try:
                k = int(comandos[1])
                aeropuertos_importantes = fly.k_aeropuertos_importantes(k)
                print(", ".join(aeropuertos_importantes))
            except Exception as e:
                Errors(e).KAeropuertoMasImportante()
            
        elif comandos[0] == "nueva_aerolinea":
            try:
                archivocsv = comandos[1]
                fly.nueva_aerolinea(archivocsv, vuelos)
                print("OK")
            except Exception as e:
                Errors(e).OptimizacionNuevaAerolinea()
                
        elif comandos[0] == "itinerario":
            try:
                fly.itinerario(comandos[1])
            except Exception as e:
                Errors(e).ItinerarioCultural()
                
        elif comandos[0] == "exportar_kml":
            try:
                archivokml = comandos[1]
                fly.exportar_kml(archivokml, caminos_mas_rapidos.desapilar())
                print("OK")
            except Exception as e:
                Errors(e).ExportarKML()