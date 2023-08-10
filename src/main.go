package main

import (
	"github.com/Thwani47/go-gin-gorm-todo-api/db"
	"github.com/gin-gonic/gin"
)

func init() {
	db.InitDb()
}

func main() {
	router := gin.Default()

	router.Run(":5000")
}
