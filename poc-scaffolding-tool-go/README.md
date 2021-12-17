# poc-scaffolding-tool-go

## Introducción

En esta poc se muestra como realizar un cliente de linea de comandos escrito en Go que permita la generación de código _boilerplate_ usando plantillas y la biblioteca Jennifer.

## Objetivo de diseño

*Requerido:*

- El sistema de permitir la generación de funciones lambda en Go usando una arquitectura hexagonal.
- Se debe poder:
    - configurar: 
        - nombre
        - tipo de evento que dispara la función (apigwproxy, etc)
    - agregar puertos
    - agregar adaptadores:
        - dynamodb
        - clientes SOAP
        - etc.

*Nota:* Para el alcance de esta PoC, solo se implementa el evento del apigwproxy y el adaptador dynamodb.


*Atributos de cálidad:* 

1. Portable: debe poder instalarse facilmente en diversos sistemas operativos sin llevar consigo dependencias de otras herramientas.
2. Mantenible: las plantillas báse o el código _boilerplate_ debe poder ser actualizado fácilmente sin necesidad de reinstalar el cli.
3. Extensible: debe ser fácil agregar nuevas plantillas.

*Deseable:*
Debe ser fácilmente adaptable a distinos canales (cli, ide, web, chatbot, etc).

*Nota:*
Con el fin de mantener la fácilidad solo se consideran casos de uso tipo _Greenfield_ en donde se necesita crear la estructura base desde cero.


## Diseño a alto nivel

En la siguiente imágen se muestra el diseño a alto nivel de la solución propuesta:

![diseño alto nivel](./doc/assets/diagrama-diseno-alto-nivel.png)


### Arquitectura función lambda prototipo

![arquitectura función lambda prototipo](./doc/assets/diagrama-arquitectura-funcion-lambda-prototipo.png)