package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"size:200;not null" json:"title"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Summary     string `gorm:"size:500" json:"summary"`
	Cover       string `gorm:"size:255" json:"cover"`
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignkey:UserID" json:"user"`
	Category    string `gorm:"size:50" json:"category"`
	Tags        string `gorm:"size:200" json:"tags"`
	ViewCount   uint   `gorm:"default:0" json:"view_count"`
	LikeCount   uint   `gorm:"default:0" json:"like_count"`
	IsPublished bool   `gorm:"default:false" json:"is_published"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "posts"
}
