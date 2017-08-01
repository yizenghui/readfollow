package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/yizenghui/readfollow/core/common"
	"github.com/yizenghui/readfollow/model"
	"github.com/yizenghui/sda"
)

// SeoTag seo 标签
type SeoTag struct {
	Title       string
	Description string
	Keywords    string
}

// Book 所用数据包
type Book struct {
	ID           uint
	Name         string
	Chapter      string
	URL          string
	BookURL      string
	Posted       string
	UpdatedAt    time.Time
	UnFollowBtm  bool
	UnFollowLink string
	FollowBtm    bool
	FollowLink   string
	Author       string
	AuthorURL    string
}

// NewBook 最近更新的书籍
type NewBook struct {
	Books     []Book
	NotUpdate bool
	SeoTag
}

//GetNewBook 获取最新更新书籍
func GetNewBook(openID string) NewBook {
	data := NewBook{}
	var book model.Book
	books := book.GetNewBooks()
	if books != nil {
		for _, b := range books {
			dbo := Book{ID: b.ID, Name: b.Name, Chapter: b.Chapter, UpdatedAt: b.UpdatedAt, Author: b.Author, AuthorURL: b.AuthorURL}
			if openID != "" {
				dbo.URL = fmt.Sprintf("/s/%d?open_id=%v", b.ID, openID)
			} else {
				dbo.URL = fmt.Sprintf("/s/%d", b.ID)
			}
			dbo.Posted = common.TransformBookPosted(b.BookURL)
			dbo.BookURL = common.TransformBookURL(b.BookURL)
			data.Books = append(data.Books, dbo)
		}
	} else {
		data.NotUpdate = true
	}

	data.Title = fmt.Sprintf("跟读最新章节")
	data.Description = fmt.Sprintf("跟读最新更新小说列表")
	data.Keywords = fmt.Sprintf("%v,%v", "跟读", "最新章节")
	return data
}

// BookInfo 所用数据包
type BookInfo struct {
	UserID       uint
	OpenID       string
	Nickname     string
	Head         string
	BookID       uint
	Name         string
	Chapter      string
	Total        string
	Author       string
	Date         string
	BookURL      string
	ChapterURL   string
	AuthorURL    string
	IsVIP        bool
	Rank         float64
	UpdatedAt    time.Time
	UnFollowBtm  bool
	UnFollowLink string
	FollowBtm    bool
	FollowLink   string
	JumpURL      string
	Posted       string
	SeoTag
}

//GetBookInfo 获取书籍详细
func GetBookInfo(id int, openID string) (BookInfo, error) {

	data := BookInfo{}

	var book model.Book
	book.GetBookByID(id)
	if book.ID == 0 {
		return data, errors.New("查无此书")
	}
	data.BookID = book.ID
	data.Name = book.Name
	data.Chapter = book.Chapter
	total := common.TransformBookTotal(book.Total)
	data.Total = common.PrintBookTotal(total)
	data.Author = book.Author
	data.BookURL = book.BookURL
	data.Posted = common.TransformBookPosted(book.BookURL)
	data.ChapterURL = book.ChapterURL
	data.IsVIP = book.IsVIP
	data.UpdatedAt = book.UpdatedAt
	data.JumpURL = common.TransformBookURL(book.BookURL)
	if openID != "" {
		var user model.User
		user.GetUserByOpenID(openID)
		if user.ID == 0 {
			// return c.Render(http.StatusOK, "404", "")
		} else {
			data.UserID = user.ID
			data.OpenID = user.OpenID
			data.Nickname = user.Nickname
			data.Head = user.Head

			if user.CheckUserIsFollowBook(book) {
				data.UnFollowBtm = true
				data.UnFollowLink = fmt.Sprintf("/unfollow/%v?open_id=%v", book.ID, user.OpenID)
			} else {
				data.FollowBtm = true
				data.FollowLink = fmt.Sprintf("/follow/%v?open_id=%v", book.ID, user.OpenID)
			}
		}
	}

	data.Title = fmt.Sprintf("跟读%v最新章节%v", data.Name, data.Chapter)
	data.Description = fmt.Sprintf("%v是%v发布于%v的小说,目前最新章节%v", data.Name, data.Author, data.Posted, data.Chapter)
	data.Keywords = fmt.Sprintf("%v,%v", data.Name, data.Author)
	return data, nil
}

// FindBook 通过url获取书籍信息
/**
url 必须是三个平台合法的书页地址
*/
func FindBook(url string) (model.Book, error) {
	var book model.Book
	book.GetBookByURL(url)
	if book.ID == 0 {
		spiderBook, _ := sda.FindBookBaseByBookURL(url)
		if spiderBook.Name != "" {
			book.Name = spiderBook.Name
			book.Author = spiderBook.Author
			book.Chapter = spiderBook.Chapter
			book.Total = spiderBook.Total
			book.AuthorURL = spiderBook.AuthorURL
			book.ChapterURL = spiderBook.ChapterURL
			book.BookURL = spiderBook.BookURL
			// 获取排行分数
			book.Rank = 0
			book.Save()
			return book, nil
		}
	} else {
		return book, nil
	}
	return book, errors.New("找不到相关书籍")
}

//FindGuide 搜索引导数据
type FindGuide struct {
	OpenID string
	SeoTag
}

//FindGuideData 搜索页SEO
func FindGuideData(openID string) FindGuide {

	data := FindGuide{}
	data.OpenID = openID
	data.Title = "搜索书籍 跟读"
	data.Keywords = "跟读,搜索"
	data.Description = "搜索书籍更新信息"

	return data
}

// SearchBook 通过关键字获取书籍信息
/**
url 必须是三个平台合法的书页地址
*/
func SearchBook(q string) ([]model.Book, error) {
	var err error
	var book model.Book
	books := book.GetBookByNameOrAuthor(q)
	if books == nil {
		err = errors.New("找不到相关书籍")
	}
	return books, err
}
