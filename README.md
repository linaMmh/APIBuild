## APIBuild
------------------------------------------------------------

## Descripción de la solución

Se creó un Api con varios servicios encargados de consultar el número PI con cierto número de decimales, además se implementó una memoria caché en Redis para no tener que calcular siempre el número de decimales de un valor dado.

## Stack tecnológico implementado

Se utilizó el lenguaje Go haciendo uso de las librerías Gin para el manejo de peticiones, y Redis para la memoria caché.

## Consideraciones funcionales y/o técnicas.

se hace uso del puerto 8080
### Para desplegar localmente
Dependencias: tener Redis de forma local, ejecutar:

```shell script
make build
```

```shell script
make run
```
or
```shell script
go build -o pi-api cmd/main.go 
```
```shell script
./pi-api 
```
### Para desplegar con docker
```shell script
make up
```
or
```shell script
docker-compose up
```

### Para consulatar la api
Existen 3 servicios principales:

```GET http://localhost:8080/getPiRandom?input_number=100```

El servicio genera un numero random con el valor enviado, y como resultado se obtiene el valor de PI con el numero de decimales del numero random.```
```
{
    "param": 100,
    "random": 97,
    "PiCalc": "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211709"
}
```

```GET http://localhost:8080/getPi?random_number=10```

El servicio toma el valor enviado y como resultado se obtiene el valor de PI con el número de decimales del valor enviado
```
{
    "random": 100,
    "PiCalc": "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798"
}
```

```DELETE http://localhost:8080/deletePi?random_number=100```

El servicio busca si existe el valor de un incide en Redis y lo elimina```

Para más información, consultar la colección de PostMan adjunta en  ```APIBuild.postman_collection.json```
##Estructura de directorios.

El api esta dividida en:

    .
    ├── ApiBuild: carpeta raiz              
    ├── cmd: directorio para el archivo main.go
    ├── common: archivos comunes utilizados en varias partes del proyecto
    ├── v1: contenedor de los casos de uso y repositorios del proyecto


##parámetros de configuración del microservicio.
Se configuraron 3 variables de entorno para el funcionamiento del proyecto

```REDIS_LOCAL_URL``` : variable de la conexion a redis

```MAX_RANDOM_PRECISION_DEFAULT``` : número máximo de decimales permitidos

```REDIS_ENABLED``` : Indica si se usara redis

##Seguridad de la API

###¿Qué componentes usarías para securitizar tu API?.

Haría una autenticación por medio de un api key, crearía otro microservicio que se encargue únicamente de validar la correcta autenticación y sería llamado antes de iniciar el proceso principal.

### ¿Cómo asegurarías tu API desde el ciclo de vida de desarrollo?
Si en la API están trabajando varios desarrolladores simultáneamente en un ambiente dockerizado se crearía un documento docker-compose.override.yml para evitar la sobrescritura de la configuración base, también se protegería el desarrollo con el uso de Gil, creando diferentes ramas para cada funcionalidad.