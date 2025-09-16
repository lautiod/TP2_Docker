# Decisiones del TP 


- **1.** Elección de la aplicación y tecnología utilizada.

- **2.** Elección de imagen base y justificación.
  
- **3.** Elección de base de datos y justificación.
  
- **4.** Estructura y justificación del Dockerfile.
  
- **5.** Configuración de QA y PROD (variables de entorno).
  
- **6.** Estrategia de persistencia de datos (volúmenes).
  
- **7.** Estrategia de versionado y publicación.
  
- **8.** Evidencia de funcionamiento: capturas de pantalla o logs mostrando:
  
**->** La aplicación corriendo en ambos entornos

**->** Conexión exitosa a la base de datos

**->** Datos persistiendo entre reinicios de contenedor

- **9.** Problemas y soluciones.
  
## 1) Elección de la aplicación y tecnología utilizada

Elegimos una aplicación web simple de usuarios (crear y listar) y decidimos desarrollarla desde cero para entender mejor cada parte del proceso y refrescar conocimientos.

Seleccionamos estas tecnologías porque ya las conocíamos y las usamos en materias anteriores, lo que nos permitió enfocarnos en la containerización más que en aprender un stack nuevo.

 - **Frontend:** React.

 - **Backend:** Go.

 - **Base de datos:** MySQL.

La carpeta del proyecto tiene dos partes: Frontend/ y Backend/.

<img width="307" height="168" alt="image" src="https://github.com/user-attachments/assets/4da39cc5-bd84-4b34-97ee-fc6397fb45f0" />

Levantamos dos entornos: QA y PROD. Ambos usan las mismas imágenes, pero se comportan distinto gracias a variables de entorno. Cada entorno tiene su propia base de datos y su propia red para que no se mezclen.

## 2) Elección de imagen base y por qué

1. **Backend:** partimos de golang:1.23-alpine para compilar y correr el servidor.

2. **Frontend:** usamos node:20-alpine para instalar y arrancar Vite.

3. **Base de datos:** mysql:8.0 (imagen oficial, estable y conocida).

_-alpine_ indica que la imagen está basada en Alpine Linux: una distribución ultraliviana pensada para que los contenedores arranquen más rápido.

## 3) Elección de base de datos y por qué

Elegimos MySQL porque es la base que más conocemos y ya hemos trabajado con ella en materias anteriores. Eso nos permitió enfocarnos en Docker y no en aprender una DB nueva.

**PROs de MySQL:**

- Tiene imagen oficial en Docker que funciona bien.

- Permite cargar scripts de inicio fácilmente (carpeta /docker-entrypoint-initdb.d/).

- Es simple de persistir: sus datos viven en /var/lib/mysql, donde montamos los volúmenes.

**Cómo la usamos:**

- Dos contenedores para los entornos de QA y PROD, cada uno con su propia base de datos y volumen. 

- Un init.sql para crear la base/tablas al primer arranque.

## 4) Estructura del Dockerfile y justificación

La idea fue tener imágenes simples y repetibles: que construyan igual en cualquier máquina, sin depender del entorno local.

**Backend:**
1. _FROM golang:1.22: imagen oficial que ya trae todo para compilar Go, de forma liviana gracias Alpine.

2. _WORKDIR /app:_ carpeta de trabajo clara dentro del contenedor.

3. _COPY go.mod go.sum + go mod download:_ primero copiamos dependencias para aprovechar la caché; si el código cambia pero las deps no, esta capa se reutiliza.

4. _COPY . .:_ copiamos el resto del código.

5. _go build -o server:_ generamos un binario listo para correr.

6. _EXPOSE 8080:_ documenta el puerto que usa la API.

7. _CMD ["./server"]:_ comando de arranque simple y directo.

**Frontend**
1. _node:20-alpine:_ versión liviana de Node, rápida de bajar y suficiente para el TP.

2. _WORKDIR /frontend:_ orden en el proyecto.

3. _COPY package*.json + npm install:_ instalamos dependencias primero (mejor caché).

4. _COPY . .:_ copiamos el resto del front.

5. _EXPOSE 5173:_ puerto del dev server de Vite.

6. _CMD ["npm","run","dev"]:_ levantamos Vite también en PROD a efectos del TP, para mantener todo simple.

## 5) Configuración de QA y PROD (variables de entorno)

Levantamos QA y PROD a partir de la misma imagen tanto para el frontend como para el backend.
La diferencia de comportamiento entre entornos se logra exclusivamente usando variables de entorno (no cambiamos código ni reconstruimos otra imagen).

En el docker-compose agrupamos las variables comunes en un bloque reutilizable (anchors de YAML).
Luego, cada servicio de QA y PROD hereda ese bloque y agrega/sobrescribe solo sus variables propias del entorno.
Así evitamos repetir configuración, mantenemos el archivo más corto y claro, y garantizamos que ambos entornos partan de la misma base.

Además, cada entorno corre en su propia red para mantenerlos aislados y evitar cruces accidentales.

**Variables comunes**

<img width="870" height="178" alt="image" src="https://github.com/user-attachments/assets/8ada69a8-8701-4e3e-9838-03ff4b888721" />

**Ejemplo: Especificación para Backend QA**

<img width="413" height="377" alt="image" src="https://github.com/user-attachments/assets/0a4f3759-499a-472e-8840-3b7c6aa52807" />

## 6) Estrategia de persistencia de datos (volúmenes)

Para que la información no se pierda al reiniciar o recrear contenedores, usamos volúmenes de Docker.
Cada entorno tiene su propio volumen de MySQL, montados en __/var/lib/mysql__, por lo que QA y PROD no comparten datos:

- QA → db_qa_data 

- PROD → db_prod_data

Además, cargamos un script de inicio (init.sql) en __/docker-entrypoint-initdb.d/.__
Ese script se ejecuta solo la primera vez que se crea el volumen (sirve para crear la base/tablas o datos de ejemplo).
Si el contenedor se reinicia y el volumen ya existe, los datos quedan tal cual.

## 7) Estrategia de versionado y publicación

Buscamos tener una forma simple de distinguir lo que está “en prueba” de lo que consideramos “estable”..

Las estiquetas que usamos fueron:

- :dev → para QA (build en prueba; se puede actualizar seguido).

- v1.0.0 → para PROD (corte estable; no se sobreescribe).

Primero construimos la imagen una vez y la publicamos en Docker Hub con ambas etiquetas. Y en docker-compose:

-> QA usa :dev.

-> PROD usa v1.0.0.

Dónde publicamos:

- Backend: danteol/tp2-backend:{dev | v1.0.0}

- Frontend: danteol/tp2-frontend:{dev | v1.0.0}

<img width="262" height="190" alt="image" src="https://github.com/user-attachments/assets/3e0d456d-2b64-41c5-b5db-4b66793b6db5" />


## 8) Evidencia de funcionamiento

**Cada uno de los frontends**

<img width="516" height="855" alt="image" src="https://github.com/user-attachments/assets/a3cbbce4-7931-4f0f-8b00-7cd3866eb9d5" />

<img width="384" height="851" alt="image" src="https://github.com/user-attachments/assets/c9277a35-2eb0-4821-ade8-928983298c84" />


**Contenedores** 

<img width="991" height="495" alt="image" src="https://github.com/user-attachments/assets/3ca2658a-b6d2-4021-a897-6e0d0dd91c3a" />


**Volumenes**

Lista de volúmenes con el comando: docker volume ls

<img width="411" height="50" alt="image" src="https://github.com/user-attachments/assets/60ae355f-8cc0-4219-b4c5-7dd359ea2ddb" />


**Algunas request**

<img width="1183" height="161" alt="image" src="https://github.com/user-attachments/assets/e1be012d-8d06-41d0-97da-5fe2e0fadb23" />

## 9) Problemas y soluciones

-> La app a veces arrancaba antes que la DB.
Solución: agregamos un chequeo de salud a MySQL y configuramos que el backend espere a que la DB esté lista.

-> Evitar mezclar QA y PROD.
Solución: usamos redes separadas (una para QA y otra para PROD).

-> Evitar repetir mucha configuración.
Solución: reutilizamos bloques de variables comunes en el docker-compose (para no copiar/pegar lo mismo en cada servicio).
