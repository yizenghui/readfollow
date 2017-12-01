package repository

import "github.com/yizenghui/sda/wechat"

// SpiderGet wechat.Article
func SpiderGet(url string) (article wechat.Article, err error) {

	article, err = wechat.Find(url)
	return
}
