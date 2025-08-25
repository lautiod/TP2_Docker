package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func mustConnect() *sql.DB {
	host := env("DB_HOST", "db")
	port := env("DB_PORT", "3306")
	user := env("DB_USER", "root")
	pass := env("DB_PASS", "root")
	name := env("DB_NAME", "testdb")
	params := env("DB_PARAMS", "parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pass, host, port, name, params)

	var database *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		database, err = sql.Open("mysql", dsn)
		if err == nil && database.Ping() == nil {
			database.SetMaxOpenConns(10)
			database.SetMaxIdleConns(5)
			database.SetConnMaxLifetime(30 * time.Minute)
			return database
		}
		time.Sleep(3 * time.Second)
	}
	log.Fatal("No se pudo conectar a la DB después de reintentos: ", err)
	return nil
}

func main() {
	// GIN_MODE se puede setear por env: release | debug | test
	ginMode := env("GIN_MODE", "release")
	gin.SetMode(ginMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	db = mustConnect()

	// Frontend
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})
	r.Static("/static", "./frontend")

	// Health
	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "down", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "up"})
	})

	// API
	r.POST("/users", func(c *gin.Context) {
		var req struct {
			Name string `json:"name"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}
		name := strings.TrimSpace(req.Name)
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre no puede estar vacío"})
			return
		}
		if _, err := db.Exec("INSERT INTO users (name) VALUES (?)", name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "usuario insertado"})
	})

	r.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name FROM users ORDER BY id DESC")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []map[string]any
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, map[string]any{"id": id, "name": name})
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	// Puerto HTTP por env
	port := env("APP_PORT", "8080")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
