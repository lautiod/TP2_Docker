# Instrucciones para:

- **1.** Construir las imágenes
- **2.** Ejecutar los contenedores
- **3.** Acceder a la aplicación (URLs, puertos)
- **4.** Conectarse a la base de datos
- **5.** Verificar que todo funciona correctamente

## 1. Construir imágenes

**Ejecutar los siguientes comandos**

_Dentro de TP2_Docker/backend_

> docker build -t danteol/tp2-backend:dev .

> docker build -t danteol/tp2-backend:v1.0.0 .

_Dentro de TP2_Docker/frontend_

> docker build -t danteol/tp2-frontend:dev .

> docker build -t danteol/tp2-frontend:v1.0.0 .

**Opcional: publicar en DockerHub**

> docker login

> docker push danteol/tp2-backend:dev

> docker push danteol/tp2-backend:v1.0.0

> docker push danteol/tp2-frontend:dev

> docker push danteol/tp2-frontend:v1.0.0

## 2. Ejecutar contenedores

**Levantar todo**

> docker compose pull              __trae las imágenes si están en Docker Hub__

> docker compose up -d             __levanta QA y PROD + bases de datos__

> docker compose ps                __ver estado (db-qa / db-prod deberían verse healthy)__

**Parar y Levantar de nuevo**

> docker compose down              __apaga todo (no borra volúmenes)__

> docker compose up -d

## 3. Acceder a la aplicación

- Frontend PROD: __http://localhost:3000/__

- Frontend QA: __http://localhost:3001/__


