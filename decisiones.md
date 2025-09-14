# Decisiones del TP 

## 1) Elección de la aplicación y tecnología utilizada

Elegimos una aplicación web simple de usuarios (crear y listar) y decidimos desarrollarla desde cero para entender mejor cada parte del proceso y refrescar conocimientos.

Seleccionamos estas tecnologías porque ya las conocíamos y las usamos en materias anteriores, lo que nos permitió enfocarnos en la containerización más que en aprender un stack nuevo.

 - **Frontend:** React + Vite.

 - **Backend:** Go con Gin.

 - **Base de datos:** MySQL.

La carpeta del proyecto tiene dos partes: Frontend/ y Backend/.

<img width="307" height="168" alt="image" src="https://github.com/user-attachments/assets/4da39cc5-bd84-4b34-97ee-fc6397fb45f0" />

Levantamos dos entornos: QA y PROD. Ambos usan las mismas imágenes, pero se comportan distinto gracias a variables de entorno. Cada entorno tiene su propia base de datos y su propia red para que no se mezclen.

## 2) Elección de imagen base y por qué

- **Backend:** partimos de golang:1.23-alpine para compilar y correr el servidor.

- **Frontend:** usamos node:20-alpine para instalar y arrancar Vite.

- **Base de datos:** mysql:8.0 (imagen oficial, estable y conocida).

_-alpine_ indica que la imagen está basada en Alpine Linux: una distribución ultraliviana pensada para que los contenedores arranquen más rápido.

## 3) Elección de base de datos y por qué

Elegimos MySQL porque es la base que más conocemos y ya hemos trabajado con ella en materias anteriores. Eso nos permitió enfocarnos en Docker y no en aprender una DB nueva.

**PROs de MySQL:**

Tiene imagen oficial en Docker que funciona bien.

Permite cargar scripts de inicio fácilmente (carpeta /docker-entrypoint-initdb.d/).

Es simple de persistir: sus datos viven en /var/lib/mysql, donde montamos los volúmenes.

**Cómo la usamos:**

Un contenedor MySQL para QA (testdbqa) y otro para PROD (testdbprod), cada uno en su red y con su volumen propio.

Un init.sql para crear la base/tablas al primer arranque.

## 4) Estructura del Dockerfile y justificación

Backend: copia módulos, instala dependencias, copia el código, compila y expone el puerto de la API.
Frontend: instala dependencias, copia el código y arranca el servidor de Vite.

La idea es tener pasos simples y repetibles para construir las imágenes sin depender de la máquina local.

## 5) Configuración de QA y PROD (variables de entorno)

La misma imagen de backend se usa en QA y PROD. Cambiamos el comportamiento con variables:

Backend QA: GIN_MODE=debug, DB testdbqa, host db-qa.

Backend PROD: GIN_MODE=release, DB testdbprod, host db-prod.

Comunes: puerto de la app, usuario/clave de la DB, etc.

El Frontend en cada entorno apunta a su backend (por ejemplo http://app-qa:8080 o http://app-prod:8080) y Vite hace de proxy para /api, así evitamos problemas de CORS.

Separarmos QA y PROD en dos redes distintas para no mezclar tráfico.

## 6) Estrategia de persistencia de datos (volúmenes)

Para que los datos no se pierdan al reiniciar contenedores:

Montamos un volumen por entorno en /var/lib/mysql.

Cargamos un script de inicio (init.sql) en /docker-entrypoint-initdb.d/ para crear la base/tablas al primer arranque.

## 7) Estrategia de versionado y publicación

Publicamos las imágenes en Docker Hub con dos etiquetas:

:dev → la usamos en QA (para probar).

v1.0.0 → la usamos en PROD (versión estable).

Para este TP, ambas etiquetas pueden apuntar a la misma imagen, así mostramos que con variables la app se comporta distinto según el entorno. En un caso real, :dev se va actualizando y v1.0.0 queda fija.

## 8) Evidencia de funcionamiento

(Acá agregaremos capturas o recortes de logs)

App corriendo en ambos entornos

Captura del frontend QA (título “Front QA”) y del frontend PROD (título “Front PROD”).

Salida de docker compose ps mostrando todos los servicios “up”.

Conexión exitosa a la base

Logs de db-qa y db-prod indicando que están ready / healthy.

Inserción y listado de usuarios desde el frontend sin errores (una captura alcanza).

Datos que persisten entre reinicios

Insertar un usuario en QA.

Reiniciar solo la base de QA: docker compose restart db-qa.

Volver a listar y mostrar que el usuario sigue ahí.

## 9) Problemas y soluciones

La app a veces arrancaba antes que la DB.
Solución: agregamos un chequeo de salud a MySQL y configuramos que el backend espere a que la DB esté lista.

Evitar mezclar QA y PROD.
Solución: usamos redes separadas (una para QA y otra para PROD).

Evitar repetir mucha configuración.
Solución: reutilizamos bloques de variables comunes en el docker-compose (para no copiar/pegar lo mismo en cada servicio).
