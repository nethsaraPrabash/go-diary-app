package model

import (
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm: "type:text;not null;" json:"content"`
	UserId uint
}