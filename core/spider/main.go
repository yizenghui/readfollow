package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yizenghui/read-follow/core/spider/qidian"
)

// Book 书籍模型
type Book struct {
	gorm.Model
	Name       string
	Chapter    string
	Total      string
	Author     string
	Date       string
	BookURL    string `sql:"index"`
	ChapterURL string
	AuthorURL  string
	IsVIP      bool
}

var db *gorm.DB

func main() {

	var err error
	db, err = gorm.Open("sqlite3", "book.db")
	// db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=spider sslmode=disable password=123456")

	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移模式
	db.AutoMigrate(&Book{})
	defer db.Close()

	syncUpdateList()

}

func syncUpdateList() {
	url := "http://a.qidian.com/?orderId=5&page=1&style=2"
	ticker := time.NewTicker(time.Minute * 2)
	for _ = range ticker.C {
		fmt.Printf("ticked at %v spider %v \n", time.Now(), url)
		go spiderBookList(url)
	}
}

func spiderBookList(url string) {
	rows, err := qidian.GetUpdateRows(url)
	if err == nil {
		for _, info := range rows {
			time.Sleep(1 * time.Second)
			syncBook(info)
		}
	}
}

// 同步职位
func syncBook(info qidian.UpdateItem) {

	var book Book
	db.Where(Book{BookURL: info.BookURL}).FirstOrCreate(&book)

	//TODO 需要验证地址是否会改变
	// 章节地址与数据库中的不同
	if book.ChapterURL != info.ChapterURL {
		book.Name = info.Name
		book.Chapter = info.Chapter
		book.ChapterURL = info.ChapterURL
		book.Author = info.Author
		book.AuthorURL = info.AuthorURL
		book.BookURL = info.BookURL
		book.IsVIP = info.IsVIP
		db.Save(&book)
		fmt.Printf("%v  %v  %v\n", book.ID, book.Name, book.Chapter)
	}

}
