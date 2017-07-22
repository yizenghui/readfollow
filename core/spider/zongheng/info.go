package zongheng

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/spider/code"
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

	html, _ := g.Html()

	item.BookURL = url
	item.ChapterURL, _ = g.Find(".update").Find(".cont").Find("a").Attr("href")

	chapterContent, _ := g.Find(".update").Find(".cont").Find("a").Html()

	chapterName := strings.TrimSpace(code.FindString(`：(?P<chapter>[^<]+)<p>`, chapterContent, "chapter"))
	if chapterName != "" {
		item.Chapter = chapterName
	} else {
		item.Chapter = g.Find(".update").Find(".cont").Find("a").Text()
	}

	dateText := strings.TrimSpace(g.Find(".update").Find(".uptime").Text())

	date := strings.TrimSpace(code.FindString(`·(?P<date>[^<]+)·`, dateText, "date"))
	if date != "" {
		item.Date = date
	} else {
		item.Date = dateText
	}

	//
	// 书名
	item.Name = code.FindString(`<meta name="og:novel:book_name" content="(?P<name>[^"]+)"/>`, html, "name")
	// 作者
	item.Author = code.FindString(`<meta name="og:novel:author" content="(?P<author>[^"]+)"/>`, html, "author")
	// 作者
	item.AuthorURL = code.FindString(`<meta name="og:novel:author_link" content="(?P<author_url>[^"]+)"/>`, html, "author_url")

	item.Total = code.FindString(`<em>·</em>字数：<span title="(?P<t>\d+)字">(?P<total>\d+)</span>字`, html, "total")

	checkIsVIP, _ := regexp.MatchString(`<em class="(?P<vip>\w+)" title="VIP作品"></em>`, html)
	if checkIsVIP {
		item.IsVIP = true
	} else {
		item.IsVIP = false
	}
	return item, nil
}
