package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hprose/hprose-golang/rpc"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/yizenghui/sda"
	"github.com/yizenghui/sda/data"
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
	PublishAt  int64 `sql:"index" default:"0"`
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

	go PostTask()
	// 启动就执行一次
	SpiderBookJobTask()
	syncUpdateList()
}

func syncUpdateList() {
	ticker1 := time.NewTicker(time.Minute * 6)
	for _ = range ticker1.C {
		SpiderBookJobTask()
	}

}

//SpiderBookJobTask 执行采集任务
func SpiderBookJobTask() {
	go SpiderBookJob("http://a.qidian.com/?orderId=5&page=%d&style=2", 20, time.Second*15)
	// go spiderBookList("http://a.qidian.com/?orderId=5&page=1&style=2")
	go SpiderBookJob("http://book.zongheng.com/store/c0/c0/b0/u0/p%d/v9/s9/t0/ALL.html", 20, time.Second*15)
	// go spiderBookList("http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/ALL.html")
	go SpiderBookJob("http://all.17k.com/lib/book/2_0_0_0_0_0_0_0_%d.html", 20, time.Second*15)
	// go spiderBookList("http://all.17k.com/lib/book/2_0_0_0_0_0_0_0_1.html")
}

// SpiderBookJob 采集
/**
FormatURL 列表格式
page 要采集到第几页
d 中止间隔
*/
func SpiderBookJob(FormatURL string, page int, d time.Duration) {
	for i := 1; i < page; i++ {
		time.Sleep(time.Second * 2)
		spiderBookList(fmt.Sprintf(FormatURL, i))
	}
}

func spiderBookList(url string) {
	rows, err := sda.GetUpdateBookByListURL(url)
	if err == nil {
		for _, info := range rows {
			time.Sleep(1 * time.Second)
			syncBook(info)
		}
	}
}

// 同步职位
func syncBook(info data.Book) {

	var book Book
	db.Where(Book{BookURL: info.BookURL}).FirstOrCreate(&book)

	// 章节地址与数据库中的不同
	if book.ChapterURL != info.ChapterURL {
		book.Name = info.Name
		book.Chapter = info.Chapter
		book.ChapterURL = info.ChapterURL
		book.Author = info.Author
		book.AuthorURL = info.AuthorURL
		book.BookURL = info.BookURL
		book.IsVIP = info.IsVIP
		book.Total = info.Total
		db.Save(&book)
		fmt.Printf("sp: %v  %v  %v\n", book.ID, book.Name, book.Chapter)
	}

}

//PostTask 同步任务
func PostTask() {
	ticker := time.NewTicker(time.Millisecond * 500)
	for _ = range ticker.C {
		go Publish()
	}
}

//Stub rpc 服务器提供接口
type Stub struct {
	Save      func(string) (string, error)
	AsyncSave func(func(string, error), string) `name:"Save"`
}

// Publish 发布
func Publish() {
	var book Book
	db.Where("publish_at = 0").First(&book)
	if book.ID > 0 {
		book.PublishAt = time.Now().Unix()
		db.Save(&book)

		if book.Chapter != "" && book.ChapterURL != "" {
			client := rpc.NewClient("http://47.92.130.14:819/")
			var stub *Stub
			client.UseService(&stub)
			postBook := TransformBook(book)
			if jsonStr, err := json.Marshal(postBook); err == nil {
				s, err := stub.Save(string(jsonStr))
				client.Close()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("ss:", s)
				}
			}
		}
	}
}

// PostBook 提交转换的数据结构
type PostBook struct {
	Name       string `json:"name"`        // 地区
	Chapter    string `json:"chapter"`     // 最小月薪
	ChapterURL string `json:"chapter_url"` // 最大月薪
	Author     string `json:"author"`      // 最大月薪
	AuthorURL  string `json:"author_url"`  // 学历
	BookURL    string `json:"book_url"`    // 工作经验
	Total      string `json:"total"`       // string默认长度为255, 使用这种tag重设。
	IsVIP      bool   `json:"is_vip"`      // string默认长度为255, 使用这种tag重设。
}

// TransformBook 数据转换
func TransformBook(book Book) PostBook {
	var pb PostBook
	pb.Name = book.Name
	pb.Chapter = book.Chapter
	pb.ChapterURL = book.ChapterURL
	pb.Author = book.Author
	pb.AuthorURL = book.AuthorURL
	pb.BookURL = book.BookURL
	pb.Total = book.Total
	pb.IsVIP = book.IsVIP
	return pb
}
