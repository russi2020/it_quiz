package models

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Title   string `json:"title" ;sql:"title"`
	Text    string `json:"text" ;sql:"text"`
	ThemeId uint   `json:"theme_id" ;sql:"theme_id" ;gorm:"foreignKey:ThemeRefer"`
	UserId  uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

// Answer To do
type Answer struct {
}

type Category struct {
	gorm.Model
	Topic  string `json:"topic" ;sql:"topic"`
	UserId uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}

type Theme struct {
	gorm.Model
	Title      string `json:"title" ;sql:"title"`
	CategoryId uint   `json:"category_id" ;sql:"category_id" ;gorm:"foreignKey:CategoryRefer"`
	UserId     uint   `json:"user_id" ;sql:"user_id" ;gorm:"foreignKey:UserRefer"`
}
