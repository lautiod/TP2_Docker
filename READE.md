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

> docker compose pull              # trae las imágenes si están en Docker Hub
> docker compose up -d             # levanta QA y PROD + bases de datos
> docker compose ps                # ver estado (db-qa / db-prod deberían verse healthy)


