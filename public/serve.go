package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
	"github.com/hprose/hprose-golang/rpc"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/yizenghui/readfollow/conf"
	"github.com/yizenghui/readfollow/model"
	"github.com/yizenghui/readfollow/repository"
)

//Template 模板
type Template struct {
	templates *template.Template
}

//Render 模板
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//Hello test
func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

//Jump 页面跳转(详细页)
func Jump(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	openID := c.QueryParam("open_id")
	data, err := repository.GetBookInfo(id, openID)
	if err != nil {
		return c.Redirect(http.StatusFound, "/404.html")
	}
	return c.Render(http.StatusOK, "jump", data)
}

//Unfollow 取消关注
func Unfollow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	openID := c.QueryParam("open_id")
	url, _ := repository.UnfollowBook(id, openID)
	return c.Redirect(http.StatusFound, url)
}

//Follow 关注
func Follow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	openID := c.QueryParam("open_id")
	url, _ := repository.FollowBook(id, openID)
	return c.Redirect(http.StatusFound, url)
}

//New 新更新的
func New(c echo.Context) error {
	openID := c.QueryParam("open_id")
	data := repository.GetNewBook(openID)
	return c.Render(http.StatusOK, "new", data)
}

//Hot 新更新的
func Hot(c echo.Context) error {
	openID := c.QueryParam("open_id")
	data := repository.GetHotBook(openID)
	return c.Render(http.StatusOK, "hot", data)
}

//User 关注
func User(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	openID := c.QueryParam("open_id")
	data, err := repository.GetUser(id, openID)
	if err != nil {
		return c.Redirect(http.StatusFound, "/404.html")
	}
	return c.Render(http.StatusOK, "user", data)
}

//Find 查找Book资源
func Find(c echo.Context) error {
	openID := c.QueryParam("open_id")
	query := c.QueryParam("q")
	data := repository.GetFind(query, openID)
	return c.Render(http.StatusOK, "find", data)
}

//Home 查找Book资源
func Home(c echo.Context) error {
	openID := c.QueryParam("open_id")
	return c.Render(http.StatusOK, "home", openID)
}

//Search 搜索本地book
func Search(c echo.Context) error {
	query := c.QueryParam("q")
	return c.Render(http.StatusOK, "search", query)
}

//PageNotFound 页面找不到
func PageNotFound(c echo.Context) error {
	return c.Render(http.StatusOK, "404", "")
}

// 接入微信接口服务
func echoWxCallbackHandler(c echo.Context) error {
	repository.WechatServe(c.Response().Writer, c.Request())
	var err error
	return err
}

//Rss sign
func Rss(c echo.Context) error {

	now := time.Now()
	feed := &feeds.Feed{
		Title:       "jmoiron.net blog",
		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&feeds.Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}

	// return c.XML(http.StatusOK, feed)

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationXMLCharsetUTF8)

	fmt.Fprintf(c.Response().Writer, rss)

	// var err error
	return err
	// c.Response().WriteHeader(http.StatusOK)
	// return xml.NewEncoder(c.Response()).Encode(rss)
}

//Sign sign
func Sign(c echo.Context) error {
	callback := c.QueryParam("callback")
	data, err := repository.CreateWebGetSignTask(callback)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(data)
	return c.Render(http.StatusOK, "sign", data)
}

// CheckSign Check Sign
func CheckSign(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	val, err := repository.GetWebGetSignTaskValue(int32(id))
	if err == nil {
		repository.RemoveSignTask(int32(id))
	}
	return c.JSON(http.StatusOK, val)
}

// SaveSign test...
func SaveSign(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	openID := c.QueryParam("open_id")
	repository.SetTaskValue(int32(id), openID)
	data, _ := repository.GetWebGetSignTaskValue(int32(id))
	return c.JSON(http.StatusOK, data)
}

func init() {
	conf.InitConfig("../conf/conf.toml")
	model.DB().AutoMigrate(&model.Book{})
}

func main() {

	// 开启rpc同步任务
	// go repository.RPCServeStart(":819")

	t := &Template{
		templates: template.Must(template.ParseGlob("weui/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/static", "./assets")
	e.File("/favicon.ico", "images/favicon.ico")
	e.Static("/images", "./images")
	service := rpc.NewHTTPService()
	service.AddFunction("Save", repository.SynchroSave)
	e.Any("/rpc", echo.WrapHandler(service))

	// e.GET("/", Home)
	e.GET("/u/:id", User)
	// e.GET("/jump/:id", Jump)
	e.GET("/s/:id", Jump)
	e.GET("/follow/:id", Follow)
	e.GET("/unfollow/:id", Unfollow)
	e.GET("/search", Search)
	e.GET("/sign", Sign)
	e.GET("/savesign", SaveSign) //测试写入签名

	e.GET("/validate", CheckSign)
	e.GET("/find", Find)
	e.GET("/hello", Hello)
	// e.GET("/hot", Hello)
	e.GET("/new", New)
	e.GET("/hot", Hot)
	e.GET("/404.html", PageNotFound)
	e.GET("/rss.xml", Rss)

	e.Any("/wx_callback", echoWxCallbackHandler)
	// Route => handler

	e.GET("/", Hot)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "域名备案中")
	// })

	// Start server 通过匹配来控制开启的端口
	e.Logger.Fatal(e.Start(conf.Conf.App.Port))
}
