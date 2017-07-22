package data

import "github.com/jinzhu/gorm"

type (

	// PostBook 转换提交的数据结构 与 RequstBook 一样的
	PostBook struct {
		Name       string `json:"name"`        // 地区
		Chapter    string `json:"chapter"`     // 最小月薪
		ChapterURL string `json:"chapter_url"` // 最大月薪
		Author     string `json:"author"`      // 最大月薪
		AuthorURL  string `json:"author_url"`  // 学历
		BookURL    string `json:"book_url"`    // 工作经验
		Total      string `json:"total"`       // string默认长度为255, 使用这种tag重设。
		IsVIP      bool   `json:"is_vip"`      // string默认长度为255, 使用这种tag重设。
	}

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
