# Proyecto de Información Geográfica

Este proyecto proporciona funcionalidades para obtener información sobre direcciones IP, países y monedas a través de APIs externas. Incluye manejo de errores, almacenamiento en caché y registro de estadísticas.

## Tabla de Contenidos

- [Características](#características)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Instalación](#instalación)
- [Uso](#uso)
- [Contribución](#contribución)
- [Licencia](#licencia)

## Características

- Obtención de información geográfica basada en direcciones IP.
- Recuperación de datos de países, incluyendo idiomas y monedas.
- Consulta de tasas de cambio de monedas.
- Registro de estadísticas de invocaciones y distancias.


## Estructura del Proyecto


```
.
├── cms
│   └── command.go             # Ejecucion inicial para validar la entrada, arranque de servicios y ejecucion
├── interfaces
│   ├── secrets.go             # Interfaz que define el comportamiento de la obtencion de los secretos
│   ├── services.go            # Interfaz que define el comportamiento de la obtencion de informacion
│   └── store.go               # Interfaz que define la forma de almacenamiento de los datos
├── models
│   ├── countryapi.go          # Definicion de la estructura de la respuesta del servicio de region
│   ├── currencyapi.go         # Definicion de la estructura de la respuesta del servicio de monedas
│   ├── errors.go              # Definicion de los errores customizados para la aplicacion
│   ├── ipapi.go               # Definicion de la estructura de la respuesta del servicio de la ip
│   ├── response.go            # Definicion de la estructura de la respuesta del proceso 'traceip'
│   └── stats.go               # Definicion de la estructura de entrada y salida para la obtencion de estadisticas
├── services
│   ├── awssecrets.go          # Implementacion del manejo de los secretos
│   ├── datastore.go           # Implementacion de la implementacion de almacenamiento (capa de persistencia)
│   ├── information.go         # Implementacion de la logica de la obtencion de la informacion
│   └── stats.go               # Logica para la obtencion, formateo y calculo de estadisticas
├── utils
│   ├── log.go                 # Configuracion dellog
│   └── utils.go               # Funciones transversales y definicion de constantes
├── main.go                    # Punto de entrada del programa
└── go.mod                     # Archivo de módulos de Go
```


## Instalación

### Requisitos

- Go (versión 1.16 o superior)
- AWS SDK para Go

### Pasos

1. Clona el repositorio:
   git clone https://github.com/tu-usuario/tu-repositorio.git
   cd tu-repositorio

2. Instala las dependencias:
   go mod tidy

3. Configura tus credenciales de AWS para acceder a Secrets Manager si es necesario.

## Uso

Para ejecutar el programa, utiliza:

go run main.go

Asegúrate de que el código esté configurado para manejar las solicitudes adecuadas según la implementación.

### Use en docker

Para poder construir un contenedor con esta aplicacion es necesario que tengas configurado docker
y algun administrador de contenedores como docker-desktop

y ejecutar los siguientes comandos en la raiz donde se encuentra el archivo 'Dockerfile'

1. Crear la imagen a partir de la configuracion del Dockerfile y el codigo fuente:
 
docker build -t service_fraud .  

2. Crear un contenedor y ejecutar a partir de la imagen generada del paso anterior, Es importante indicar que la obtencion de las claves para el consumo de las api se realizan por medio en este caso se secrets manager por lo que es necesario configurar un usuario con permisos limitados para poder acceder a este servicio

docker run -it -e AWS_ACCESS_KEY_ID=valor1 -e AWS_SECRET_ACCESS_KEY=valor2 service_fraud

## Contribución

Si deseas contribuir al proyecto:

1. Haz un fork del repositorio.
2. Crea una rama para tu característica.
3. Realiza tus cambios y haz un commit.
4. Haz un push a tu rama.
5. Abre un Pull Request.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

