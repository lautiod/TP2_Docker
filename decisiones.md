# Decisiones del TP (Docker)

## 1) Elección de la aplicación y tecnología utilizada

Hicimos una app sencilla de usuarios para practicar la containerización:

Frontend: React + Vite (rápido para probar).

Backend: Go con Gin (liviano y fácil de ejecutar).

Base de datos: MySQL.

La carpeta del proyecto tiene dos partes: Frontend/ y Backend/.

<img width="307" height="168" alt="image" src="https://github.com/user-attachments/assets/4da39cc5-bd84-4b34-97ee-fc6397fb45f0" />


## 2) Elección de imagen base y por qué

Backend: partimos de golang:1.22 para compilar y correr el servidor.

Frontend: usamos node:20-alpine para instalar y arrancar Vite.

Base de datos: mysql:8.0 (imagen oficial, estable y conocida).

(Nota: para un “prod real” podríamos achicar las imágenes, pero para el TP priorizamos simplicidad.)

## 3) Elección de base de datos y por qué

Elegimos MySQL porque:

Es muy usada y tiene imagen oficial fácil de levantar.

Soporta scripts de inicio (para crear tablas/datos al arrancar).

Su carpeta de datos es clara (/var/lib/mysql), lo que facilita persistir.

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
