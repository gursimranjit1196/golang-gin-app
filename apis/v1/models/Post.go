package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Name    string `gorm:"size:255;not null;unique" json:"name" binding:"required"`
	Content string `gorm:"size:255;not null" json:"content" binding:"required"`
	UserID  int    `json:"user_id"`

	User User
}

func (p *Post) CreatePost(DB *gorm.DB) (*Post, error) {
	var err error
	err = DB.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return p, err
	}

	return p, nil
}
