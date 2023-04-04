package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"yasotubez/models"

	// import models from models folder
	"os"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	/* 	err := godotenv.Load()
	   	if err != nil {
	   		panic(err)
	   	} */

	// Database connection
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println(dbURL)
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

	r.GET("/result", func(c *gin.Context) {
		// get tasks from database
		rows, err := db.Query("SELECT * FROM tasks")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		var tasks []models.Task
		for rows.Next() {
			var task models.Task
			err := rows.Scan(&task.ID, &task.Deadline, &task.Title, &task.Description)
			if err != nil {
				panic(err)
			}
			tasks = append(tasks, task)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

		// reverse the tasks array
		for i := len(tasks)/2 - 1; i >= 0; i-- {
			opp := len(tasks) - 1 - i
			tasks[i], tasks[opp] = tasks[opp], tasks[i]
		}

		// get done tasks from database
		rows, err = db.Query("SELECT * FROM done")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		var doneTasks []models.Task
		for rows.Next() {
			var task models.Task
			err := rows.Scan(&task.ID, &task.Deadline, &task.Title, &task.Description)
			if err != nil {
				panic(err)
			}
			doneTasks = append(doneTasks, task)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

		// reverse the doneTasks array
		for i := len(doneTasks)/2 - 1; i >= 0; i-- {
			opp := len(doneTasks) - 1 - i
			doneTasks[i], doneTasks[opp] = doneTasks[opp], doneTasks[i]
		}

		c.HTML(http.StatusOK, "result.tmpl", gin.H{
			"tasks": tasks,
			"done":  doneTasks,
		})
	})

	r.POST("/result", func(c *gin.Context) {
		deadline := c.PostForm("deadline")
		title := c.PostForm("title")
		desc := c.PostForm("desc")

		_, err := db.Exec("INSERT INTO tasks (deadline, title, description) VALUES ($1, $2, $3)", deadline, title, desc)
		if err != nil {
			panic(err)
		}
		// redirect to result get
		c.Redirect(http.StatusMovedPermanently, "/result")
	})

	r.POST("/add-done", func(c *gin.Context) {
		id := c.PostForm("id")
		// move the task from tasks table to done table
		_, err := db.Exec("INSERT INTO done SELECT * FROM tasks WHERE id = $1", id)
		if err != nil {
			panic(err)
		}
		// delete the task from tasks table
		_, err = db.Exec("DELETE FROM tasks WHERE id = $1", id)
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/result")
	})

	r.POST("/delete", func(c *gin.Context) {
		id := c.PostForm("id")
		// delete the task from tasks table
		_, err := db.Exec("DELETE FROM done WHERE id = $1", id)
		if err != nil {
			panic(err)
		}
		c.Redirect(http.StatusMovedPermanently, "/result")
	})

	r.Run() // listen and serve on
}
