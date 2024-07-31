#!/usr/bin/python3.
from grafo import Grafo
from math import inf
from collections import deque
from heap import Heap
from cola import Cola

#+-------------------------------------------------------------------------------------------------+#
#+----------------------------------+CAMINOS MINIMOS+----------------------------------------------+#
#+-------------------------------------------------------------------------------------------------+#
def camino_minimo_dijkstra(grafo, origen, destino = None):
    """
    Aplica dijkstra desde origen, devuelve diccionario de padres y distancias
    """
    distancia, padre = {}, {}
    for v in grafo.obtener_vertices():
        distancia[v] = inf
    distancia[origen] = 0
    padre[origen] = None
    q = Heap()
    q.encolar(origen, 0)
    while not q.esta_vacia():
        v = q.desencolar()
        if destino != None and v == destino:
            return padre, distancia
        
        for w in grafo.adyacentes(v):
            
            if (distancia[v] + float(grafo.peso_arista(v, w)) < distancia[w]):
                distancia[w] = distancia[v] + float(grafo.peso_arista(v, w))
                padre[w] = v
                q.encolar(w, distancia[w])
    return padre, distancia 

def camino_minimo_bfs(grafo, origen):
    """
    Aplica bfs desde origen, devuelve diccionario de padres y distancias
    """
    distancia, padre, visitado = {}, {}, {}
    for v in grafo.obtener_vertices():
        distancia[v] = inf
    distancia[origen] = 0
    padre[origen] = None
    visitado[origen] = True
    q = Cola()
    q.encolar(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if (w not in visitado):
                distancia[w] = distancia[v] + 1
                padre[w] = v
                visitado[w] = True
                q.encolar(w)
    return padre, distancia

#+-------------------------------------------------------------------------------------------------+#
#+------------------------------ARBOLES DE TENDIDO MINIMO (MST)+-----------------------------------+#
#+-------------------------------------------------------------------------------------------------+#

def mst_prim(grafo):
    """
    Devuelve un arbol de tendido minimo
    """
    v = grafo.vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    q = Heap()
    for w in grafo.adyacentes(v):
        q.encolar((v, w), grafo.peso_arista(v, w))
    arbol = Grafo(es_dirigido= False, vertices= grafo.obtener_vertices())
    while not q.esta_vacia():
        (v, w) = q.desencolar()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, grafo.peso_arista(v, w))
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if x not in visitados:
                q.encolar((w, x), grafo.peso_arista(w, x))
    return arbol

#+-------------------------------------------------------------------------------------------------+#
#+------------------------------------+CENTRALIDAD+------------------------------------------------+#
#+-------------------------------------------------------------------------------------------------+#

def ordenar_vertices(grafo, distancias):
    """
    Ordena los vertices de mayor a menor distancia
    """
    vertices = grafo.obtener_vertices()
    vertices_ordenados = sorted(vertices, key=lambda v: distancias[v])
    return vertices_ordenados

def centralidad(grafo):
    """
    Devuelve la centralidad (tipo Betweenes) que tiene cada vertice en un grafo, en funcion de cuantas veces aparece en los 
    caminos mínimos entre todos los pares de vértices (sin ser uno de los extremos)
    """
    cent = {}
    for v in grafo.obtener_vertices(): 
        cent[v] = 0
    for v in grafo.obtener_vertices():
        padre, distancias = camino_minimo_dijkstra(grafo, v)
        cent_aux = {}
        for w in grafo.obtener_vertices(): 
            cent_aux[w] = 0
        vertices_ordenados = ordenar_vertices(grafo, distancias)
        for w in vertices_ordenados:
            if padre[w] is None: continue
            cent_aux[padre[w]] += 1 + cent_aux[w]
        for w in grafo.obtener_vertices():
            if w == v: continue
            cent[w] += cent_aux[w]
    return cent

#+-------------------------------------------------------------------------------------------------+#
#+------------------------------------+ORDEN TOPOLOGICO+-------------------------------------------+#
#+-------------------------------------------------------------------------------------------------+#

def grados_entrada(grafo):
    """
    Devuelve los grados de entrada de todos los vertices en una lista
    """
    gr_entrada = {}
    for vertice in grafo.obtener_vertices():
        gr_entrada[vertice] = 0
    
    for vertice in grafo.obtener_vertices():
        for ady in grafo.adyacentes(vertice):
            gr_entrada[ady] = gr_entrada[ady] + 1
    
    return gr_entrada

def topologico_grados(grafo):
    """
    Devuelve una lista de los vertices a recorrer del grafo en funcion de un ordenamiento topologico (grados)
    """
    g_ent = grados_entrada(grafo)
    q = Cola()
    for v in grafo.obtener_vertices():
        if g_ent[v] == 0:
            q.encolar(v)
    resultado = []
    while not q.esta_vacia():
        v = q.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.encolar(w)
    return resultado
