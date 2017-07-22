package qidian

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// QiDianUpdateList 起点小说更新列表
var QiDianUpdateList = "http://a.qidian.com/?orderId=5&page=1&style=2"

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
	g.Find(".rank-table-list tbody tr").Each(func(i int, content *goquery.Selection) {
		// 书详细页
		item.BookURL, _ = content.Find(".name").Attr("href")
		item.ChapterURL, _ = content.Find(".chapter").Attr("href")
		// 书名
		item.Name = strings.TrimSpace(content.Find(".name").Text())
		// 章节
		item.Chapter = strings.TrimSpace(content.Find(".chapter").Text())
		// 作者
		item.Author = strings.TrimSpace(content.Find(".author").Text())
		// 作者详细页
		item.AuthorURL, _ = content.Find(".author").Attr("href")
		// 小说更新时间
		item.Date = strings.TrimSpace(content.Find(".date").Text())
		// 字数
		item.Total = strings.TrimSpace(content.Find(".total").Text())

		checkLinkIsJobInfo, _ := regexp.MatchString(`vip(?P<reader>\w+).qidian.com`, item.ChapterURL)
		if checkLinkIsJobInfo {
			item.IsVIP = true
		} else {
			item.IsVIP = false
		}

		rows = append(rows, item)
	})

	return rows, nil
}
