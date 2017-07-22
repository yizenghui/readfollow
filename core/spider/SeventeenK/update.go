package SeventeenK

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SeventeenKUpdateList 纵横小说网(男生站)
var SeventeenKUpdateList = "http://all.17k.com/lib/book.html"

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
	g.Find("table tbody tr").Each(func(i int, content *goquery.Selection) {
		// 书名
		item.Name = strings.TrimSpace(content.Find(".td3").Find(".jt").Text())
		// tr有空行
		if item.Name != "xxxx" {

			// 书籍地址
			item.BookURL, _ = content.Find(".td3").Find(".jt").Attr("href")
			// 章节
			item.Chapter = strings.TrimSpace(content.Find(".td4").Find("a").Eq(0).Text())
			// 章节地址
			item.ChapterURL, _ = content.Find(".td4").Find("a").Attr("href")

			// 作者名
			item.Author = strings.TrimSpace(content.Find(".td6").Text())
			// 作者详细页
			item.AuthorURL, _ = content.Find(".td6").Find("a").Attr("href")

			// 字数
			item.Total = strings.TrimSpace(content.Find(".td5").Text())

			// 更新时间
			item.Date = strings.TrimSpace(content.Find(".td7").Text())

			checkIsVIP, _ := content.Find(".td4").Find(".vip").Attr("title")
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
