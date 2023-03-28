package model

import (
	"diary_api/database"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (ent *Entry) Save() (*Entry, error) {
	err := database.Database.Create(&ent).Error
	if err != nil {
		return &Entry{}, nil
	}

	return ent, nil
}
