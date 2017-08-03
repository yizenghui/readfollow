package model

import "time"

// Book 数据模型
type Book struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time  `sql:"index"`
	DeletedAt  *time.Time `sql:"index"`
	Name       string     `gorm:"size:255"`
	Chapter    string     `gorm:"size:255"`
	Total      string     `gorm:"size:255"`
	Author     string     `gorm:"size:255"`
	Date       string     `gorm:"size:255"`
	BookURL    string     `sql:"index"`
	ChapterURL string     `gorm:"size:255"`
	AuthorURL  string     `gorm:"size:255"`
	IsVIP      bool
	Rank       float64 `sql:"index"`
	Users      []User  `gorm:"many2many:user_books;"` // 关注书的用户
}

// GetBookByID 获取用户关注的书籍
func (book *Book) GetBookByID(id int) {
	DB().First(book, id)
}

//GetFollowUser 获取书籍的所有关注者
func (book *Book) GetFollowUser() []User {
	var users []User
	DB().Model(&book).Association("users").Find(&users)
	return users
}

// GetNewBooks 获取最新100本书
func (book *Book) GetNewBooks() []Book {
	var books []Book
	DB().Limit(100).Order("updated_at Desc").Find(&books)
	return books
}

// GetHotBooks 获取最热100本书
func (book *Book) GetHotBooks() []Book {
	var books []Book
	DB().Limit(100).Order("rank Desc").Find(&books)
	return books
}

// GetBookByURL 通过url获取书籍信息 如果没有的话进行初始化 注：没有地址相同的两本书
func (book *Book) GetBookByURL(url string) {
	DB().Where(Book{BookURL: url}).FirstOrInit(book)
}

// GetBookByName 通过书名获取书籍信息
func (book *Book) GetBookByName(name string) (books []Book) {
	DB().Where(Book{Name: name}).Find(&books)
	return
}

// GetBookByAuthor 通过作者获取书籍信息
func (book *Book) GetBookByAuthor(author string) (books []Book) {
	DB().Where(Book{Author: author}).Find(&books)
	return
}

// GetBookByNameOrAuthor 通过书名或作者名获取书籍信息
func (book *Book) GetBookByNameOrAuthor(query string) (books []Book) {
	// TODO 这个后期再优化，做成tag搜索
	DB().Where(Book{Name: query}).Or(Book{Author: query}).Find(&books)
	return
}

// Save 保存用户信息
func (book *Book) Save() {
	DB().Save(&book)
}
