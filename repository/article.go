package repository

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/GanEasy/wxrankapi/orm"
	"github.com/yizenghui/sda/wechat"
)

// ArticleItem 返回前台使用的文章信息
type ArticleItem struct {
	ID            uint
	Title         string
	Author        string
	Cover         string
	PubAt         time.Time
	MediaTagID    int
	MediaTagTitle string
	Like          int
	Hate          int
	Rank          float64
	URL           string
}

//GetArticleData 获取文章数据，转义返回
func GetArticleData(article orm.Article) (data ArticleItem) {

	ArticleAfter(&article)

	mediaTagID := 0
	mediaTagTitle := ""
	for _, t := range article.Tags {
		tagMsg, _ := GetTagMsg(int(t))
		if tagMsg.Type == string("wxid") {
			mediaTagID = int(tagMsg.ID)
			mediaTagTitle = tagMsg.Title
		}
	}

	data = ArticleItem{
		ID:            article.ID,
		Title:         article.Title,
		Author:        article.Author,
		Cover:         article.Cover,
		PubAt:         article.PubAt,
		MediaTagID:    mediaTagID,
		MediaTagTitle: mediaTagTitle,
		Like:          int(article.Like),
		Hate:          int(article.Hate),
		Rank:          article.Rank,
		URL:           article.URL,
	}
	return
}

//GetArticlesData 获取文章数据，转义返回
func GetArticlesData(articles []orm.Article) (rows []ArticleItem) {

	for _, article := range articles {
		rows = append(rows, GetArticleData(article))
	}
	return
}

func init() {
	orm.DB().AutoMigrate(&orm.Tag{}, &orm.Article{}, &orm.Media{}, &orm.Post{})
}

// Find wechat.Article
func Find(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}

func update(url string) (article wechat.Article, err error) {
	article, err = wechat.Find(url)
	return
}

// Insert wechat.Article
func Insert(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}

//GetArticle 获取文章列表
func GetArticle(limit, offset, tag int) (articles []orm.Article, err error) {
	var a orm.Article
	// var articles []orm.Article

	// articles = a.GetArticle(limit, offset, tag, "rank DESC,pub_at DESC,id ASC")
	articles = a.GetArticle(limit, offset, tag, "pub_at DESC,id ASC")

	// orm.DB().Offset(offset).Limit(limit).Order("rank DESC").Find(&articles)
	for key, article := range articles {
		ArticleAfter(&article)
		articles[key] = article
	}

	return
}

//GetArticleCursorByID 通过ID游标方式获取最新收录文章
func GetArticleCursorByID(id, limit int, tags []int64) (rows []ArticleItem, err error) {
	var a orm.Article
	articles := a.GetArticleCursorByID(id, limit, tags)
	rows = GetArticlesData(articles)
	return
}

//GetArticleCursorByRank 通过Rank游标方式获取热门文章
func GetArticleCursorByRank(rank float64, limit int, tags []int64) (rows []ArticleItem, err error) {
	var a orm.Article
	articles := a.GetArticleCursorByRank(rank, limit, tags)
	rows = GetArticlesData(articles)
	return
}

// ArticleAfter 修改文章
func ArticleAfter(article *orm.Article) {
	article.Cover = "http://pic3.readfollow.com/" + base64.URLEncoding.EncodeToString([]byte(article.Cover))
	// article.Cover = "http://localhost:8005/file/" + base64.URLEncoding.EncodeToString([]byte(article.Cover))
	// article.Cover = "https://readfollow.com/file/" + base64.URLEncoding.EncodeToString([]byte(article.Cover))
	article.URL = strings.Replace(article.URL, `http://`, `https://`, -1)
	article.URL = strings.Replace(article.URL, `#rd`, "&scene=27#wechat_redirect", 1)

	article.Title = strings.Replace(article.Title, `\x26quot;`, `"`, -1)
	article.Title = strings.Replace(article.Title, `\x26amp;`, `&`, -1)
	article.Title = strings.Replace(article.Title, `\x26gt;`, `>`, -1)
	article.Title = strings.Replace(article.Title, `\x0a`, "\n", -1)

	article.Intro = strings.Replace(article.Intro, `\x0a`, "\n", -1)
	article.Intro = strings.Replace(article.Intro, `\x26quot;`, `"`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26gt;`, `>`, -1)
	article.Intro = strings.Replace(article.Intro, `\x26amp;`, `&`, -1)

}

//View ..
func View(id int) (a orm.Article, err error) {

	// var a orm.Article
	a.GetArticleByID(id)

	if a.Title == "" {
		err = errors.New("内容异常")
		return
	}

	a.View++

	if a.ID != 0 {
		a.Rank = Rank(int(a.View), 0, a.PubAt.Unix())
		a.Save()
	}

	ArticleAfter(&a)
	return
}

//Like ..
func Like(id int) (a orm.Article, err error) {

	// var a orm.Article
	a.GetArticleByID(id)

	if a.Title == "" {
		err = errors.New("内容异常")
		return
	}

	a.Like++

	if a.ID != 0 {
		a.Rank = ArticleRank(a)
		a.Save()
	}

	ArticleAfter(&a)

	return
}

//Hate ..
func Hate(id int) (a orm.Article, err error) {

	// var a orm.Article
	a.GetArticleByID(int(id))

	if a.Title == "" {
		err = errors.New("内容异常")
		return
	}

	a.Hate++

	if a.ID != 0 {
		a.Rank = ArticleRank(a)
		a.Save()
	}

	ArticleAfter(&a)

	return
}

// Post ... url
func Post(url string) (err error) {
	var post orm.Post
	post.GetPostByURL(url)
	// if post.State == 0 { // 检查提交状态
	var a orm.Article
	article, err := wechat.Find(url)
	if err == nil {

		if article.URL == "" {
			return errors.New("不支持该链接！")
		}

		media, err := GetMediaByAppID(article.AppID)
		if err != nil {
			return errors.New("公众号信息出错！")
		}
		// 如果公众号是新收录的
		if media.State == 0 {
			media.AppName = article.AppName
			media.Cover = article.RoundHead
			media.State = 1

			// 公众号ID作为一个标签
			var tag orm.Tag
			tag.GetTagByName(article.AppID)
			if tag.Type == "" {
				tag.Type = "wxid"
				tag.Title = article.AppName
				// tag.IsShow = 0
				tag.Save()
			}

			media.Tags = append(media.Tags, int64(tag.ID))

			media.Save()
		}

		post.ArticleURL = article.URL
		post.State = 1
		post.Save()
		a.GetArticleByURL(article.URL)
		a.MediaID = media.ID
		a.Title = article.Title
		a.Intro = article.Intro
		a.Cover = article.Cover
		a.Author = article.Author
		// todo 标签管理，需要保留自定义标签
		a.Tags = media.Tags // 文章的标签等于公众号的标签
		a.View++
		i64, err := strconv.ParseInt(article.PubAt, 10, 64)
		if err != nil {
			// fmt.Println(err)
			return errors.New("时间转化失败")
		}
		a.PubAt = time.Unix(i64, 0)

		a.Rank = ArticleRank(a)

		// panic(a.ID)

		a.Save()
		// fmt.Println(a)
	}
	// }
	return
}

//ArticleRank 获取文章RANK 该rank具有维一性，可作游标  通过创建时间计算期分值，以赞同否定作为偏移值
func ArticleRank(article orm.Article) (rank float64) {
	// 根据创建时间进行排列
	t := article.CreatedAt.Unix()

	// 如果否定的票数比赞同多5票，将会得0分，多6票的话，将得负分
	r := Rank(int(article.Like+5), int(article.Hate), t)

	// 得分截取 0.12 分位
	s := fmt.Sprintf("%.2f", r)

	// 发布时间截取 日、时
	h := article.PubAt.Format("0215")

	//todo 截取文章ID最多6位
	// 组装成 rank 值   0.12 得分 + 日、时 + 文章ID
	str := fmt.Sprintf(`%v%v%d`, s, h, article.ID)

	// 字符串格式化成浮点数
	rank, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return
	}
	return
}

// Remove wechat.Article
func Remove(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}
