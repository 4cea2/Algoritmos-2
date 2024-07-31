#!/usr/bin/python3
import heapq

class Heap:
    def __init__(self, is_min_heap=True):
        self.heap = []
        self.contador = 0
        self.multiplier = 1 if is_min_heap else -1

    def encolar(self, elemento, prioridad):
        heapq.heappush(self.heap, (self.multiplier * prioridad, self.contador, elemento))
        self.contador += 1

    def desencolar(self):
        return heapq.heappop(self.heap)[2]

    def esta_vacia(self):
        return len(self.heap) == 0
    
    def ver_maximo(self):
        if self.esta_vacia():
            raise AssertionError("El heap se encuentra vacío")
        else:
            return self.heap[0]
        
    def heapify(self, diccionario):
        #ordena por valor y si con iguales desempata con clave(minimo)
        self.heap = [(self.multiplier * valor, self.contador, clave) for clave, valor in diccionario.items()]
        self.contador = len(diccionario)
        heapq.heapify(self.heap)
