package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username" binding:"required"`
	Email    string `gorm:"size:100;not null;unique" json:"email" binding:"required"`
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Create(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (p *User) GetAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

func (u *User) GetUser(db *gorm.DB, id uint64) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", id).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Where("id = ?", u.ID).Updates(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}
