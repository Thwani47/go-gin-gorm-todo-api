package models

import (
	"errors"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `gorm:"not null" json:"title"`
	Done  bool   `gorm:"default:false"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	if todo.Title == "" {
		err = errors.New("Title is required")
	}
	return
}
