// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/GanEasy/wxrankapi/orm"
	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/sda/wechat"
)

func Test_Post(t *testing.T) {
	// err := Post("https://mp.weixin.qq.com/s/e_0BhJ0vOvIqt610MjgKLA")
	// if err != nil {
	// 	t.Error(err)
	// }
	// err = Post("https://mp.weixin.qq.com/s/lu9yU_Wvd8yQs6alZ7OBqg")
	// if err != nil {
	// 	t.Error(err)
	// }
	// err = Post("https://mp.weixin.qq.com/s/widHAGN7zOymivIUihkuTw")
	// if err != nil {
	// 	t.Error(err)
	// }
	// err = Post("https://mp.weixin.qq.com/s/wXOhpZHdzlHlf5jgC54AaA")
	// if err != nil {
	// 	t.Error(err)
	// }
	err := Post("https://mp.weixin.qq.com/s?__biz=MjM5MDE0Mjc4MA==&mid=2650999105&idx=3&sn=6de522819d357a2bf50242d933b4abbe&chksm=bdbef1128ac97804d8f5528267b67c8510fbf698f6561520e38bd2d276f7b6e7dbd5d7f788f6&mpshare=1&scene=1&srcid=111664iGQ9EBw9bURidnhSjh#rd")
	if err != nil {
		t.Error(err)
	}
}

func Test_Pos22t(t *testing.T) {
	url := "https://mp.weixin.qq.com/s/e_0BhJ0vOvIqt610MjgKLA"
	var post orm.Post
	post.GetPostByURL(url)
	var a orm.Article
	article, err := wechat.Find(url)
	if err == nil {

		if article.URL == "" {
			t.Error("url err")
		}

		media, err := GetMediaByAppID(article.AppID)
		if err != nil {
			t.Error(err)

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
		a.Tags = media.Tags // 文章的标签等于公众号的标签

		i64, err := strconv.ParseInt(article.PubAt, 10, 64)
		if err != nil {
			// fmt.Println(err)

			t.Error(err)
		}
		a.PubAt = time.Unix(i64, 0)
		a.Save()
		// panic(a.ID)
		t.Error(a)
	} else {
		t.Error(err)
	}

}

func Test_GetArticle(t *testing.T) {
	articles, _ := GetArticle(5, 0, 0)
	fmt.Println(articles)
	t.Error(articles)

}

func Test_GetArticleRank(t *testing.T) {
	articles, _ := GetArticle(5, 200, 0)

	// fmt.Println(articles)
	for _, a := range articles {
		fmt.Println(a)
		log.Fatal(ArticleRank(a))

	}

	// t.Error(articles)

}
func Test_GetArticleCursorByID(t *testing.T) {

	tags := []int64{1}
	articles, _ := GetArticleCursorByID(500, 10, tags)

	// fmt.Println(articles)
	// for _, a := range articles {
	// 	// fmt.Println(ArticleRank(a))

	// }

	t.Error(articles)

}
func Test_GetArticleCursorByRank(t *testing.T) {

	tags := []int64{}
	articles, _ := GetArticleCursorByRank(61.3837037037037, 10, tags)

	// fmt.Println(articles)
	// for _, a := range articles {
	// 	// fmt.Println(ArticleRank(a))

	// }

	t.Error(articles)

}
func Test_PostLink(t *testing.T) {
	// link := url.QueryEscape("https://mp.weixin.qq.com/s?__biz=MjM5NzgzNTUyNA==&mid=2650373560&idx=1&sn=84964bd5ce47084ae175806a4bf279da&chksm=bede3bd389a9b2c540644ba0161a5ece8f5c84b3088da9eccc3d8d0712e21cf31f6a637b4940#rd")
	link := url.QueryEscape("http://mp.weixin.qq.com/s?__biz=MzI4ODU0Nzk0NA==&amp;mid=2247484254&amp;idx=2&amp;sn=83cec43c784e699a827a3c9987c9171c&amp;chksm=ec3df659db4a7f4fa457a95daa77a07534e3172902ca27f39ed7a91d9170d260073ace202325&amp;mpshare=1&amp;scene=1&amp;srcid=1010Yx7oukwdLbpQ9QY31fwU#rd")
	doc, err := goquery.NewDocument(fmt.Sprintf("http://localhost:8888/fetch?url=%v", link))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Html())
}
