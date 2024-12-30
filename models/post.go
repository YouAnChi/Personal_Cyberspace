package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"size:200;not null" json:"title"`
	Content     string `gorm:"type:text;not null" json:"content"`
	HTMLContent string `gorm:"type:text" json:"html_content"`
	Summary     string `gorm:"size:500" json:"summary"`
	Cover       string `gorm:"size:255" json:"cover"`
	Author      string `gorm:"size:100" json:"author"`
	Category    string `gorm:"size:50" json:"category"`
	Tags        string `gorm:"size:200" json:"tags"`
	ViewCount   uint   `gorm:"default:0" json:"view_count"`
	LikeCount   uint   `gorm:"default:0" json:"like_count"`
	IsPublished bool   `gorm:"default:true" json:"is_published"`
}

// TableName 指定表名
func (Post) TableName() string {
	return "posts"
}
