package common

import "github.com/yizenghui/readfollow/model"

// func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

// ByUpdate 通过更新时间排序
type ByUpdate []model.Book

func (a ByUpdate) Len() int           { return len(a) }
func (a ByUpdate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdate) Less(i, j int) bool { return a[j].UpdatedAt.Before(a[i].UpdatedAt) }

var newBooks []model.Book

//GetNewBooks 获取100本最近更新的书籍
func GetNewBooks() []model.Book {
	if newBooks == nil {
		var book model.Book
		newBooks = book.GetNewBooks()
	}
	newBooks = newBooks[0:100]
	return newBooks
}

// AddNewBook 添加一本刚更新的书籍
func AddNewBook(book model.Book) {

}
