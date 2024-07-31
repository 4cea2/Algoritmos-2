# TDAs

- [**Pila Dinámica**](https://github.com/4cea2/Algoritmos-2/tree/main/tdas/pila): Estructura LIFO que usa un slice redimensionable para almacenar elementos. Permite apilar y desapilar de manera eficiente.

- [**Cola Enlazada**](https://github.com/4cea2/Algoritmos-2/tree/main/tdas/cola): Estructura FIFO basada en nodos enlazados, con punteros al primer y último nodo para gestionar los elementos.

- [**Lista Enlazada**](https://github.com/4cea2/Algoritmos-2/tree/main/tdas/lista): Estructura que facilita la inserción y eliminación en cualquier parte, usando nodos conectados por punteros y acompañado por un iterador que recorre la lista permitiendo insertar y eliminar en cualquier parte.

- [**Diccionario con Hash**](https://github.com/4cea2/Algoritmos-2/tree/main/tdas/diccionario): Emplea una tabla de hash abierto para asociar claves con valores, ofreciendo acceso rápido a los datos. Tiene su propio iterador.

- [**Diccionario con ABB**](https://github.com/4cea2/Algoritmos-2/tree/main/tdas/diccionario): Usa un árbol binario de búsqueda para almacenar pares clave-valor, optimizando búsqueda, inserción y eliminación. Cuenta con iteradores, inclusive de rangos.

- [**Heap**](https://github.com/4cea2/Algoritmos-2/tree/main/tdas/cola_prioridad): Estructura en forma de árbol binario completo que puede ser un max-heap o min-heap, manteniendo una propiedad específica en los nodos.

Todas las estructuras son genéricas, permitiendo el almacenamiento de cualquier tipo de dato, y cada una incluye sus propios tests.

# TPs

- **TP0**: Introducción al lenguaje de programación Go y repaso de conceptos clave de la programación.

- **TP1**: Implementación de un sistema de votación con detección de fraudes. El sistema recibe comandos, los procesa y muestra los resultados. Utiliza los TDAs previamente implementados para gestionar los votos y verificar irregularidades.

- **TP2**: Desarrollo de un sistema para gestionar la entrada y salida de aviones en un aeropuerto. El sistema permite ordenar, filtrar, analizar y obtener información sobre los vuelos. Emplea los TDA implementados, adaptándolos a las complejidades y requisitos específicos.

- **TP3**: Programa para analizar las conexiones que hay entre aeropuertos en funcion de grafos. Implementa un **TDA Grafo** para aplicar algoritmos relacionados con caminos y tendidos minimos, centralidad y de ordenamientos topologicos.


---

Tanto los TDAs de lista, diccionario y heap, como los TPs 1, 2 y 3, fueron realizados durante la cursada en conjunto con  [Fuentes Acuña, Brian Alex](https://github.com/Alex-f98).
