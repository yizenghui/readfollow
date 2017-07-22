package qidian

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Info 58PC职位列参数
type Info struct {
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

//GetInfo 获取书籍详细
func GetInfo(url string) (Info, error) {

	var item Info
	g, e := goquery.NewDocument(url)
	if e != nil {
		return item, e
	}

	item.BookURL = url
	item.ChapterURL, _ = g.Find(".book-info-detail").Find(".update").Find(".cf").Find(".blue").Attr("href")

	item.Chapter = strings.TrimSpace(g.Find(".book-info-detail").Find(".update").Find(".cf").Find(".blue").Text())

	item.Date = strings.TrimSpace(g.Find(".book-info-detail").Find(".update").Find(".cf").Find(".time").Text())

	item.Author = strings.TrimSpace(g.Find(".book-info").Find(".writer").Text())

	item.AuthorURL, _ = g.Find(".book-info").Find(".writer").Attr("href")

	item.Total = strings.TrimSpace(g.Find(".book-info").Find("p").Eq(2).Find("em").Eq(0).Text())

	checkLinkIsJobInfo, _ := regexp.MatchString(`vip(?P<reader>\w+).qidian.com`, item.ChapterURL)
	if checkLinkIsJobInfo {
		item.IsVIP = true
	} else {
		item.IsVIP = false
	}
	return item, nil
}
