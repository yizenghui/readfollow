package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/yizenghui/readfollow/conf"
	"github.com/yizenghui/readfollow/repository"
	"github.com/yizenghui/sda/code"
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

	url := code.ExplainBookDetailedAddress(query)
	if url != "" {
		book, err := repository.FindBook(url)
		if err == nil {
			return c.Redirect(http.StatusFound, fmt.Sprintf("/s/%d?open_id=%v", book.ID, openID))
		}
		return c.Render(http.StatusOK, "hello", "找不到您所想要的资源")
	} else {
		// 通过书名搜索
		if query != "" {
			books, err := repository.SearchBook(query)
			if err == nil {
				// 取第一本
				book := books[0]
				return c.Redirect(http.StatusFound, fmt.Sprintf("/s/%d?open_id=%v", book.ID, openID))
			}
		}
	}
	return c.Render(http.StatusOK, "find", openID)

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

func init() {
	conf.InitConfig("../conf/conf.toml")
}

func main() {

	// 开启rpc同步任务
	go repository.RPCServeStart(":819")

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	// e.Static("/static", "../assets")

	// e.GET("/", Home)
	e.GET("/u/:id", User)
	// e.GET("/jump/:id", Jump)
	e.GET("/s/:id", Jump)
	e.GET("/follow/:id", Follow)
	e.GET("/unfollow/:id", Unfollow)
	e.GET("/search", Search)
	e.GET("/find", Find)
	e.GET("/hello", Hello)
	// e.GET("/hot", Hello)
	e.GET("/new", New)
	e.GET("/404.html", PageNotFound)

	e.Any("/wx_callback", echoWxCallbackHandler)
	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "域名备案中")
	})

	// Start server 通过匹配来控制开启的端口
	e.Logger.Fatal(e.Start(conf.Conf.App.Port))
}
