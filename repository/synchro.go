package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/yizenghui/readfollow/core/common"
	"github.com/yizenghui/readfollow/core/event"
	"github.com/yizenghui/readfollow/model"
)

type (

	//RequstBook POST 请求参数获取
	RequstBook struct {
		Name       string `json:"name" valid:"Required; MaxSize(24)"`
		Chapter    string `json:"chapter" valid:"Required; MaxSize(64)"`
		Total      string `json:"total" valid:"MaxSize(24);"`
		Author     string `json:"author" valid:"Required; MaxSize(24);"`
		BookURL    string `json:"book_url" valid:"Required; MaxSize(255);"`
		ChapterURL string `json:"Chapter_url" valid:"MaxSize(255);"`
		AuthorURL  string `json:"author_url" valid:"MaxSize(255);"`
		IsVIP      bool   `json:"is_vip"`
	}
)

// 数据库对象

// RPCServeStart 开启 RPC 服务
func RPCServeStart(listen string) {
	// conf.InitConfig("../conf/conf.toml")
	service := rpc.NewHTTPService()
	service.AddFunction("Save", SynchroSave, rpc.Options{})
	http.ListenAndServe(listen, service)
}

// SynchroRequest 获取请求数据
func SynchroRequest(str string) (RequstBook, error) {
	var qbook RequstBook
	var err error
	json.Unmarshal([]byte(str), &qbook)

	valid := validation.Validation{}

	b, err := valid.Valid(&qbook)
	if err != nil {
		err = errors.New("数据验证错误")
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		err = errors.New("参数异常")
	}

	if qbook.BookURL == "" {
		err = errors.New("同步职位失败")
	}
	return qbook, err
}

//SynchroSave 同步保存数据接口
func SynchroSave(str string) string {
	// fmt.Println(str)
	var book model.Book
	qbook, err := SynchroRequest(str)
	if err != nil {
		return "err: " + err.Error() + "!"
	}
	book.GetBookByURL(qbook.BookURL)

	// 记录章节地址是否变化
	bookChapterURL := book.ChapterURL

	err = SynchroRequstBookSaveData(&book, qbook)
	if err != nil {
		return "err: " + err.Error() + "!"
	}

	fansNum := 0 // 粉丝数量
	if book.ID != 0 && bookChapterURL != book.ChapterURL {
		users := book.GetFollowUser()
		if users != nil {
			event.BookUpdateNotice(book, users)
			fansNum = len(users)
		}
	}

	// 获取排行分数
	book.Rank = common.GetRank(fansNum, 0, time.Now().Unix(), 0)
	book.Save()

	// 最近更新列表添加上本书
	AddNewBook(book)

	// fmt.Println("print ", book.ID, book.Name, book.Chapter, book.Rank)
	// bookString, _ := json.Marshal(book)
	return fmt.Sprintf("%v %v %v", book.ID, book.Name, book.Chapter)
}

//SynchroRequstBookSaveData 把请求的数据包转成数据模型中的参数
func SynchroRequstBookSaveData(book *model.Book, qb RequstBook) error {

	book.Name = qb.Name
	book.Chapter = qb.Chapter
	book.Total = qb.Total
	book.Author = qb.Author
	book.BookURL = qb.BookURL
	book.ChapterURL = qb.ChapterURL
	book.AuthorURL = qb.AuthorURL
	book.IsVIP = qb.IsVIP

	return nil
}
