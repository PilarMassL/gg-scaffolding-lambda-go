# poc-scaffolding-tool-go

## Introducción

En esta poc se muestra como realizar un cliente de linea de comandos escrito en Go que permita la generación de código _boilerplate_ usando plantillas y la biblioteca Jennifer.

## Objetivo de diseño

*Requerido:*

- El generador debe permitir la generación de funciones lambda en Go usando una arquitectura hexagonal.
- Se debe poder:
    - configurar: 
        - nombre
        - tipo de evento que dispara la función (apigwproxy, etc)
    - agregar puertos
    - agregar adaptadores:
        - repositorio a tabla dynamodb
        - clientes API REST
        - proxy a servicios web SOAP
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

Ejemplo de uso desde la terminal:
*Nota:* Esto es una visión de cómo sería la interfaz de linea de comandos.

Para iniciar un proyecto basado en funciones
```
ggen init <PROYECT_NAME>
[INFO] Generando desde la plantilla...
[INFO] proyecto base generado ✅
```

Para agregar una función, debe hacerse sobre un proyecto iniciado.
```
ggen add fn
> ¿Cuál es el nombre de la función? (my-function)
> ¿Qué tipo de evento disparará la función?
 👉 - API Proxy HTTP
    - S3
    - EventBridge

[INFO] Generando desde la plantilla...
[INFO] Función generada ✅
```

Para agregar un puerto y adaptador a una función:
```
ggen add port -f my-function
> ¿Cuál es el nombre del puerto?
> ¿Agregar adaptador? (Y/n)
> Tipo de adaptador
    - Generico
 👉 - Repositorio tabla DynamoDB
    - Reposiorio Documento DocumentDB
    - Cliente API REST
    - Proxy WS SOAP

[INFO] Generando desde la plantilla...
[INFO] Puerto y Adaptador generado ✅
```

Para generar un servicio:
```
ggen add svc -f my-function
> Tipo de servicio
    - Dominio
 👉 - Aplicación
 > Nombre del servicio (my-service)

[INFO] Generando desde la plantilla...
[INFO] Servicio generado ✅
```

## Diseño a alto nivel

En la siguiente imágen se muestra el diseño a alto nivel de la solución propuesta:

![diseño alto nivel](./doc/assets/diagrama-diseno-alto-nivel.png)

Para la creación del cliente en linea de comandos se usa [Cobra](https://github.com/spf13/cobra).


### Arquitectura función lambda prototipo

![arquitectura función lambda prototipo](./doc/assets/diagrama-arquitectura-funcion-lambda-prototipo.png)

## Decisiones tomadas de antemano

El objetivo es crear un prototipo dogmático con algunas decisiones predefinidas.

Para registros de auditoría (logs):
- [Zap](https://github.com/uber-go/zap)

gestión de trazas:
- [Jaeger](https://github.com/open-telemetry/opentelemetry-go)

Validación de entrada:
- [Validator](https://github.com/go-playground/validator)

Resiliencia
- [Histryx-go](https://github.com/afex/hystrix-go)

Para pruebas:
- Mocking: [GoMock](https://github.com/golang/mock)
- [Testify]() 

Para el consumo de servicios de AWS se utilizan las bibliotecas proporcionadas por AWS.


## Mini-tutorial: Cómo agregar un nuevo Adapter al generador.
En este tutorial se explica cómo agregar un adapter para un storga basado en S3.

TODO

## Mini-tutorial: Cómo usar el generador en una solución real.
En este mini-tutorial se explica cómo construir un bot de slack que sirva para inicalizar proyectos serverless básados en AWS lambda usando Golang con ggen.

TODO

## Referencias

1. https://github.com/RanchoCooper/go-hexagonal
2. https://blog.ralch.com/articles/golang-code-generation-tool-implementation/
3. https://github.com/dave/jennifer
4. https://awesomeopensource.com/project/Alikhll/golang-developer-roadmap

