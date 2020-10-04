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

func (p *Post) GetAllPosts(DB *gorm.DB) (*[]PostWithAssociationSerializer, error) {
	var err error
	posts := []PostWithAssociationSerializer{}
	err = DB.Debug().Table("posts").Preload("User").Find(&posts).Error

	if err != nil {
		return &[]PostWithAssociationSerializer{}, err
	}

	return &posts, nil
}
