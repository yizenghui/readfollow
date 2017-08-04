package common

import (
	"sort"

	"github.com/yizenghui/readfollow/model"
)

// func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

// ByUpdate 通过更新时间排序
type ByUpdate []model.Book

// ByRank 通过热度排序
type ByRank []model.Book

func (a ByUpdate) Len() int           { return len(a) }
func (a ByUpdate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdate) Less(i, j int) bool { return a[j].UpdatedAt.Before(a[i].UpdatedAt) }

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return a[j].Rank > a[i].Rank }

var newBooks []model.Book
var hotBooks []model.Book

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
	AppanNewBook(book)
	sort.Sort(ByUpdate(newBooks))
	newBooks = newBooks[0:100]
}

// AddHotBook 添加一本刚更新的书籍
func AddHotBook(book model.Book) {
	if book.Rank > 0 {
		AppanHotBook(book)
		sort.Sort(ByRank(hotBooks))
		if len(hotBooks) > 100 {
			hotBooks = hotBooks[0:100]
		}
	}
}

// GetHotBooks 获取热门小说 top 100
func GetHotBooks() []model.Book {
	if hotBooks == nil {
		var book model.Book
		hotBooks = book.GetHotBooks()
	}
	if len(hotBooks) > 100 {
		hotBooks = hotBooks[0:100]
	}
	return hotBooks
}

// AppanHotBook 添加一本刚更新的书籍 (过滤重复)
func AppanHotBook(book model.Book) {

	if hotBooks == nil {
		hotBooks = GetHotBooks()
	}
	temp := hotBooks
	for k, b := range hotBooks {
		if b.ID == book.ID {
			temp = append(temp[:k], temp[k+1:]...)
		}
	}
	hotBooks = temp
	hotBooks = append(hotBooks, book)

}

// AppanNewBook 添加一本刚更新的书籍 (过滤重复)
func AppanNewBook(book model.Book) {
	if newBooks == nil {
		newBooks = GetNewBooks()
	}
	temp := newBooks
	for k, b := range newBooks {
		if b.ID == book.ID {
			temp = append(temp[:k], temp[k+1:]...)
		}
	}
	newBooks = temp
	newBooks = append(newBooks, book)
}
