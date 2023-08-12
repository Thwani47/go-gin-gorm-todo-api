package db

import (
	"github.com/Thwani47/go-gin-gorm-todo-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.Todo{})

}
