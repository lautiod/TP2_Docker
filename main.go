package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error

	// user:password@tcp(host:puerto)/dbname
	dsn := "root:root@tcp(db:3306)/testdb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// probar conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la DB:", err)
	}

	r := gin.Default()

	// insertar usuario
	r.POST("/users", func(c *gin.Context) {
		var req struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}

		_, err := db.Exec("INSERT INTO users (name) VALUES (?)", req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "usuario insertado"})
	})

	// obtener usuarios
	r.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var name string
			rows.Scan(&id, &name)
			users = append(users, map[string]interface{}{
				"id":   id,
				"name": name,
			})
		}

		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}
