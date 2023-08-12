package controllers

import (
	"net/http"

	"github.com/Thwani47/go-gin-gorm-todo-api/db"
	"github.com/Thwani47/go-gin-gorm-todo-api/models"
	"github.com/gin-gonic/gin"
)

func AddTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo added successfully", "todo": todo})
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	result := db.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func GetCompleteTodos(c *gin.Context) {
	var todos []models.Todo

	result := db.DB.Where("done = ?", true).Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func GetTodo(c *gin.Context) {
	var todo models.Todo

	result := db.DB.First(&todo, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo

	result := db.DB.Delete(&todo, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo

	result := db.DB.First(&todo, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching todo"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = db.DB.Save(&todo)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "todo": todo})
}
