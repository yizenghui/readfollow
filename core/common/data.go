package common

import "github.com/jinzhu/gorm"

type (

	//RequstBook POST 请求参数获取
	RequstBook struct {
		Name       string `json:"name" valid:"Required; MaxSize(24)"`
		Chapter    string `json:"chapter" valid:"Required; MaxSize(64)"`
		Total      string `json:"total" valid:"MaxSize(24);"`
		Author     string `json:"author" valid:"Required; MaxSize(12);"`
		BookURL    string `json:"book_url" valid:"Required; MaxSize(255);"`
		ChapterURL string `json:"Chapter_url" valid:"MaxSize(255);"`
		AuthorURL  string `json:"author_url" valid:"MaxSize(255);"`
		IsVIP      bool   `json:"is_vip"`
	}

	// Book 数据模型
	Book struct {
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
	}
)

//RequstBookSaveData 把请求的数据包转成数据模型中的参数
func RequstBookSaveData(book *Book, qb RequstBook) error {

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
