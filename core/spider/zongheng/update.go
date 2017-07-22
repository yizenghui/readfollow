package zongheng

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ZongHengUpdateList 纵横小说网(男生站)
var ZongHengUpdateList = "http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/ALL.html"

// UpdateRows 58PC职位列表记录集
type UpdateRows []UpdateItem

// UpdateItem 58PC职位列参数
type UpdateItem struct {
	Name       string
	Chapter    string
	Total      string
	Author     string
	Date       string
	BookURL    string
	ChapterURL string
	AuthorURL  string
	IsVIP      bool
}

//GetUpdateRows 获取更新页列表内容
func GetUpdateRows(url string) (UpdateRows, error) {

	var rows UpdateRows
	var item UpdateItem
	g, e := goquery.NewDocument(url)
	if e != nil {
		return rows, e
	}

	// 下列内容于 2017年4月4日 20:50:24 抓取
	g.Find(".main_con li").Each(func(i int, content *goquery.Selection) {
		// 书名
		item.Name = strings.TrimSpace(content.Find(".chap").Find(".fs14").Text())
		// li有空行
		if item.Name != "" {

			// 书籍地址
			item.BookURL, _ = content.Find(".chap").Find(".fs14").Attr("href")
			// 章节
			item.Chapter = strings.TrimSpace(content.Find(".chap").Find("a").Eq(1).Text())
			// 章节地址
			item.ChapterURL, _ = content.Find(".chap").Find("a").Eq(1).Attr("href")

			// 作者名
			item.Author = strings.TrimSpace(content.Find(".author").Text())
			// 作者详细页
			item.AuthorURL, _ = content.Find(".author").Find("a").Attr("href")

			// 字数
			item.Total = strings.TrimSpace(content.Find(".number").Text())

			// 更新时间
			item.Date = strings.TrimSpace(content.Find(".time").Text())

			checkIsVIP, _ := content.Find(".chap").Find(".vip").Attr("title")
			if checkIsVIP != "" {
				item.IsVIP = true
			} else {
				item.IsVIP = false
			}

			rows = append(rows, item)
		}
	})

	return rows, nil
}
