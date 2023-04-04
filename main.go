package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
