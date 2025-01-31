package model

import (
	"gorm.io/gorm"
	"go-diary-app/database"
)

type Entry struct {
	gorm.Model
	Content string `gorm: "type:text;not null;" json:"content"`
	UserId uint
}

func (entry *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&entry).Error
	if err != nil {
		return &Entry{}, err
	}
	return entry, nil
}

