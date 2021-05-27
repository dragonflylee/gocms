package model

import (
	"time"
)

// Article 文章
type Article struct {
	ID        int64      `json:"-" gorm:"primaryKey" form:"-"`
	AdminID   int64      `json:"-" form:"-"`
	Title     string     `json:"title" gorm:"index;size:256;not null" form:"title"`
	Tags      string     `json:"tags,omitempty" gorm:"size:256" form:"tags"`
	Content   string     `json:"content,omitempty" gorm:"type:text" form:"content"`
	Remark    string     `json:"remark,omitempty" gorm:"size:512" form:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"not null" form:"-"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" form:"-"`
	Admin     *Admin     `json:"-" form:"-"`
}

// Query 获取文章
func (a *Article) Query(where ...interface{}) error {
	return db.Take(a, where...).Error
}

// Update 更新文章
func (a *Article) Update() error {
	if a.ID <= 0 {
		return db.Create(a).Error
	}
	return db.Model(a).Updates(a).Error
}

// GetArticles 获取文章列表
func GetArticles(filter ...Scope) ([]Article, error) {
	var list []Article
	tx := db.Preload("Admin")
	for _, s := range filter {
		tx = tx.Scopes(s.Scope)
	}
	err := tx.Find(&list).Error
	return list, err
}

// GetArticleNum 获取文章数量
func GetArticleNum(filter ...Scope) (int64, error) {
	var nums int64
	tx := db.Model(&Article{})
	for _, s := range filter {
		tx = tx.Scopes(s.Scope)
	}
	err := tx.Count(&nums).Error
	return nums, err
}
