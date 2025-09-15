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

## 4. Conectarese a la base de datos

No es necesario abrir el puerto 3306 al host. Podés entrar al contenedor y usar el cliente mysql que ya trae la imagen.

Para ello ejecutar los siguientes comandos abriendo una consola:

**DB de QA**
> docker exec -it tp2_docker-db-qa-1 mysql -uroot -proot

__Ya dentro de MySQL__

> SHOW DATABASES;

> USE testdbqa;

> SHOW TABLES;

> SELECT * FROM users;

**DB de PROD**
> docker exec -it tp2_docker-db-prod-1 mysql -uroot -proot

__Ya dentro de MySQL__

> SHOW DATABASES;

> USE testdbprod;

> SHOW TABLES;

> SELECT * FROM users;

## 5. Verificar que todo funciona correctamente

### La app debe estar corriendo en ambos entornos

Abrí http://localhost:3001 (QA) y http://localhost:3000 (PROD).

En cada uno, probá crear y listar usuarios.

-> En QA deberías ver el título/indicador de QA.

-> En PROD, el de PROD.

### Conexión exitoso con la base de datos

Si podés agregar usuarios desde el front y luego listarlos, la API está conectando bien con MySQL.

También podés chequear el estado de MySQL atráves del siguiente comando:

> docker ps -a    __debería decir (healthy) para db-qa y db-prod__

**Persistencia de datos:** si reiniciás o eliminás los contenedores, los datos de la base se mantienen gracias a los volúmenes.

En caso de que quieras eliminar los contenedores junto a los volumenes, ejecuta el siguiente comando:

> docker compose down -v






