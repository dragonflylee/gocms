package model

import (
	"html/template"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/russross/blackfriday.v2"
)

// Article 文章
type Article struct {
	ID        int64      `json:"-" gorm:"primary_key"`
	Title     string     `json:"title" gorm:"index;size:256;not null"`
	Tags      string     `json:"tags,omitempty" gorm:"size:256"`
	Content   string     `json:"content,omitempty" gorm:"type:text"`
	Remark    string     `json:"remark,omitempty" gorm:"size:512"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"not null"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Render 渲染
func (a Article) Render() template.HTML {
	md := []byte(strings.ReplaceAll(a.Content, "\r\n", "\n"))
	return template.HTML(blackfriday.Run(md))
}

// Query 获取文章
func (a *Article) Query() error {
	return db.New().First(a, "id = ?", a.ID).Error
}

// Update 更新文章
func (a *Article) Update() error {
	return db.New().Model(a).Updates(a).Error
}

// GetArticles 获取文章列表
func GetArticles(filter ...func(*gorm.DB) *gorm.DB) ([]Article, error) {
	var list []Article
	err := db.New().Scopes(filter...).Find(&list).Error
	return list, err
}

// GetArticleNum 获取文章数量
func GetArticleNum(filter ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var nums int64
	err := db.New().Model(&Article{}).Scopes(filter...).Count(&nums).Error
	return nums, err
}
