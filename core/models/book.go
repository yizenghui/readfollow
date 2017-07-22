package models

import "github.com/jinzhu/gorm"

// Book 数据模型
type Book struct {
	gorm.Model
	Name       string `gorm:"size:255"`
	Chapter    string `gorm:"size:255"`
	Total      string `gorm:"size:255"`
	Author     string `gorm:"size:255"`
	Date       string `gorm:"size:255"`
	BookURL    string `sql:"index"`
	ChapterURL string `gorm:"size:255"`
	AuthorURL  string `gorm:"size:255"`
	IsVIP      bool
	Rank       float64 `sql:"index"`
	Users      []User  `gorm:"many2many:user_books;"` // 关注书的用户
}

//GetFollow 获取书籍的关注者
func (book *Book) GetFollow() {
	// 这里面因为要使用 DB 的问题，暂时没满意的解决方案
}
