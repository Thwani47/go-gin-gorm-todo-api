package main

import (
	"github.com/Thwani47/go-gin-gorm-todo-api/controllers"
	"github.com/Thwani47/go-gin-gorm-todo-api/db"
	"github.com/gin-gonic/gin"
)

func init() {
	db.InitDb()
}

func main() {
	router := gin.Default()

	router.POST("/api/v1/new", controllers.AddTodo)
	router.GET("/api/v1/todos", controllers.GetTodos)
	router.GET("/api/v1/todos/complete", controllers.GetCompleteTodos)
	router.GET("/api/v1/todo/:id", controllers.GetTodo)
	router.DELETE("/api/v1/todo/:id", controllers.DeleteTodo)
	router.PUT("/api/v1/todo/:id", controllers.UpdateTodo)

	router.Run(":5000")
}
