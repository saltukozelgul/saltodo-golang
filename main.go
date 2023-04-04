package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Database connection
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/result", func(c *gin.Context) {
		deadline := c.PostForm("deadline")
		title := c.PostForm("title")
		desc := c.PostForm("desc")
		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"deadline": deadline,
			"title":    title,
			"desc":     desc,
		})
	})
	r.Run() // listen and serve on
}
