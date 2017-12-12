package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/yizenghui/reader"
	"github.com/yizenghui/readfollow/job"
	r "github.com/yizenghui/readfollow/repository"
)

type (
	//Stats 结构
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

//NewStats New Stats
func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: make(map[string]int),
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

//Articles 文章接口
func Articles(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	tag, _ := strconv.Atoi(c.QueryParam("tag"))

	if limit <= 0 || limit > 100 {
		limit = 10
	}
	// limit = 10
	if offset < 0 || offset > 500 {
		offset = 0
	}

	articles, err := r.GetArticle(limit, offset, tag)

	if err != nil {

	}

	return c.JSON(http.StatusOK, articles)
}

//NewArticles 最新收录文章接口
func NewArticles(c echo.Context) error {

	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	tags := r.Str2Int64(c.QueryParam("tag"))

	id, _ := strconv.Atoi(c.QueryParam("id"))

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	articles, _ := r.GetArticleCursorByID(id, limit, tags)

	return c.JSON(http.StatusOK, articles)
}

//HotArticles 文章接口 根据热门程序进行游标提取
func HotArticles(c echo.Context) error {

	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	tags := r.Str2Int64(c.QueryParam("tag"))

	rank, _ := strconv.ParseFloat(c.QueryParam("rank"), 64)

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	articles, _ := r.GetArticleCursorByRank(rank, limit, tags)

	return c.JSON(http.StatusOK, articles)
}

//Tags 标签列表接口
func Tags(c echo.Context) error {
	t := c.QueryParam("type")

	tags, err := r.GetTagByType(t)

	if err != nil {

	}
	return c.JSON(http.StatusOK, tags)
}

//Search 标签列表搜索接口
func Search(c echo.Context) error {
	t := c.QueryParam("s")

	tags, err := r.GetTagsByTitle(t)

	if err != nil {

	}
	return c.JSON(http.StatusOK, tags)
}

//Tag 标签详细
func Tag(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tag, err := r.Tag(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, tag)
}

//GetTagByMediaID 通过公众号ID获取标签详细
func GetTagByMediaID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tag, err := r.GetTagByMediaID(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, tag)
}

//View 阅读
func View(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	article, err := r.View(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, article)
}

//Like 喜欢
func Like(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	article, err := r.Like(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, article)
}

//Hate 讨厌
func Hate(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	article, err := r.Hate(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, article)
}

//Media 公众号
func Media(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	media, err := r.GetMediaByID(id)

	if err != nil {

	}

	return c.JSON(http.StatusOK, media)
}

//Fetch get 报料接口
func Fetch(c echo.Context) error {
	url := c.QueryParam("url")
	// fmt.Println(url)
	if url != "" {
		// r.Post(url)
		// 列队任务, 防止高并发攻击
		job.JobQueue <- job.Job{
			Task: &job.TaskSpider{
				URL: url,
			},
		}
		return c.JSON(http.StatusOK, "1")
	}
	return c.JSON(http.StatusOK, "0")
}

//JsSDK 微信JS接口
func JsSDK(c echo.Context) error {
	url := c.QueryParam("url")

	js, _ := r.GetSign(url)
	return c.JSON(http.StatusOK, js)
}

//Post 报料接口
func Post(c echo.Context) error {
	url := c.FormValue("url")
	// fmt.Println("url", url)
	if url != "" {
		err := r.Post(url)
		if err != nil {
			return c.JSON(http.StatusOK, "0")
		}
		return c.JSON(http.StatusOK, "1")
	}
	return c.JSON(http.StatusOK, "0")
}

//imgServe 图片服务接口
func imgServe(c echo.Context) error {
	input := c.Param("url")
	uDec, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		r.PrintErrorImageHandler(c.Response().Writer, c.Request())
	} else {
		r.PrintImageHandler(string(uDec), c.Response().Writer, c.Request())
	}
	var err2 error
	return err2
}

//GetContent 获取正文 临时放在这里面，做小程序测试api
func GetContent(c echo.Context) error {
	urlStr := c.QueryParam("url")

	info, err := reader.GetContent(urlStr)
	if err != nil {
		return c.JSON(http.StatusOK, "0")
	}

	input := []byte(info.Content)
	unsafe := blackfriday.MarkdownCommon(input)
	content := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	bh := fmt.Sprintf(`
			<html>
			<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
			<title>%v</title>
			<body>
			%v
			</body>
			</html>
			`, info.Title, string(content[:]))

	g, e := goquery.NewDocumentFromReader(strings.NewReader(bh))
	if e != nil {
		return c.JSON(http.StatusOK, "0")
	}
	// html := fmt.Sprintf(`<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	// 						<link rel="preload" href="https://yize.gitlab.io/css/main.css" as="style" />
	// 						%v`, string(content[:]))
	// return c.HTML(http.StatusOK, html)
	info.Content = g.Text()

	type Info struct {
		Title   string        `json:"title"`
		Content template.HTML `json:"content"`
		PubAt   string        `json:"pub_at"`
	}

	return c.JSON(http.StatusOK, Info{
		info.Title,
		template.HTML(info.Content),
		info.PubAt,
	})
}

//GetList 获取列表 临时放在这里面，做小程序测试api
func GetList(c echo.Context) error {
	urlStr := c.QueryParam("url")
	if urlStr == "" {
		return c.JSON(http.StatusOK, "0")
	}
	links, _ := reader.GetList(urlStr)
	return c.JSON(http.StatusOK, links)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	//-------------------
	// Custom middleware
	//-------------------
	// Stats
	s := NewStats()
	e.Use(s.Process)
	// 展示统计
	e.GET("/stats", s.Handle) // Endpoint to get stats

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, Welcome to api.readfollow.com !\n")
	// })

	// 请求抓取
	e.GET("/fetch", Fetch)
	e.POST("/post", Post)

	/*******以下是api请求*********/
	e.POST("/api/post", Post)
	// 获取公众号接口
	e.GET("/api/media/:id", Media)
	// 用户查看文章时请求该接口
	e.GET("/api/view/:id", View)
	// 赞同文章
	e.GET("/api/like/:id", Like)
	// 否定文章 如果否定比赞同多5票，评分为0
	e.GET("/api/hate/:id", Hate)

	// 获取微信文章接口
	e.GET("/api/article", Articles)

	// 获取微信文章接口
	e.GET("/api/new", NewArticles)
	e.GET("/api/hot", HotArticles)

	e.GET("/api/jssdk", JsSDK)

	// 获取标签接口
	e.GET("/api/tags", Tags)
	e.GET("/api/tag/:id", Tag)
	e.GET("/api/search", Search)
	e.GET("/api/gettagbymedia/:id", GetTagByMediaID)
	/********以上是api请求******/

	// 图片文件服务
	e.GET("/file/:url", imgServe)

	/***以下是兼容前端的**/
	e.File("/", "static/dist/index.html")
	e.File("/t/:id", "static/dist/index.html")
	e.File("/tags", "static/dist/index.html")
	e.File("/hot", "static/dist/index.html")
	e.File("/new", "static/dist/index.html")
	e.Static("static", "static/dist/static")
	/***以上是兼容前端的**/

	// 临时做小程序api
	e.GET("/minapp/getlist", GetList)
	e.GET("/minapp/getcontent", GetContent)

	e.File("logo.png", "images/80x80logo.png")
	e.File("favicon.ico", "images/favicon.ico")
	//e.Logger.Fatal(e.Start(":8005"))

	e.Logger.Fatal(e.StartAutoTLS(":443"))

}
