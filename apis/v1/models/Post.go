package models

import (
	"gin-app/apis/v1/config/validator"

	"github.com/jinzhu/gorm"
	"gorm.io/gorm/clause"
)

type Post struct {
	BaseModel
	Name    string `gorm:"size:255;not null;unique" json:"name" binding:"required" validate:"gte=3,lte=20"`
	Content string `gorm:"size:255;not null" json:"content" binding:"required" validate:"gte=5,lte=100"`
	UserID  int    `json:"user_id"`
	User    User   `binding:"-"`
}

func (p *Post) CreatePost(DB *gorm.DB) (*Post, error) {
	v := validator.GetValidator()
	isValid := v.Struct(*p)
	if isValid != nil {
		return nil, isValid
	}

	var err error
	err = DB.Debug().Model(&Post{}).Omit(clause.Associations).Create(&p).Error
	if err != nil {
		return p, err
	}

	return p, nil
}

func (p *Post) GetAllPosts(DB *gorm.DB) (*[]Post, error) {
	var err error
	posts := []Post{}
	err = DB.Debug().Model(&Post{}).Preload("User").Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}

	return &posts, nil
}
