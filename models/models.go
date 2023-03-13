package models

import "gorm.io/gorm"

type Todolist struct {
	gorm.Model
	Name string
	Description string 
}