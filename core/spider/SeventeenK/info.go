package SeventeenK

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

	// 页面上的更新时间
	dateText := strings.TrimSpace(code.FindString(`<em>更新：(?P<date>[^<]+)</em>`, html, "date"))

	vipChapterMap := code.SelectString(`<span class="time">更新时间：(?P<time>[^<]+)</span>(?P<_>\s+)最新vip章节：<a(?P<_>\s+)href="(?P<url>[^"]+)"(?P<_>\s+)target="_blank">(?P<chapter>[^<]+)</a>`, html)

	chapterMap := code.SelectString(`<span class="time">更新时间：(?P<time>[^<]+)</span>(?P<_>\s+)最新免费章节：<a(?P<_>\s+)href="(?P<url>[^"]+)"(?P<_>\s+)target="_blank">(?P<chapter>[^<]+)</a>`, html)

	// 普通章节更新时间需要与页面更新时间一致
	if t, ok := chapterMap["time"]; ok && t == dateText {

		item.Chapter, _ = chapterMap["chapter"]
		item.ChapterURL, _ = chapterMap["url"]
		item.Date = t
		// fmt.Println(chapterMap)
	} else if t, ok := vipChapterMap["time"]; ok { // && t == dateText
		//  && t == dateText

		item.Chapter, _ = vipChapterMap["chapter"]
		item.ChapterURL, _ = vipChapterMap["url"]
		item.Date = t
		// fmt.Println(vipChapterMap)
	}

	//
	// 书名
	item.Name = strings.TrimSpace(g.Find("h1").Find("a").Text())
	// 作者
	item.Author = strings.TrimSpace(g.Find(".AuthorInfo").Find(".name").Text())
	// 作者页
	item.AuthorURL, _ = g.Find(".AuthorInfo").Find(".name").Attr("href")

	// 字数
	item.Total = strings.TrimSpace(g.Find(".BookData").Find(".red").Text())

	// 如果含有VIP章节字眼，为VIP作品
	checkIsVIP, _ := regexp.MatchString(`最新v(?P<vip>\w+)章节：`, html)
	if checkIsVIP {
		item.IsVIP = true
	} else {
		item.IsVIP = false
	}
	return item, nil
}
