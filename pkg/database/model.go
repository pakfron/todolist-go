package database

import (
	"gorm.io/gorm"
)

type Todolist struct {
	gorm.Model
	ID    int `gorm:"AUTO_INCREMENT"`
	Todo  string
	Check bool `gorm:"default:false"`
	IsDel bool `gorm:"default:false"`
}
