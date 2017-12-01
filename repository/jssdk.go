package repository

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/jssdk"
)

// Post has and belongs to many languages, use `Post_languages` as join table
type JsSign struct {
	URL       string `json:"url" `
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
	NonceStr  string `json:"nonceStr"`
}

//GetSign 获取排名
func GetSign(url string) (js JsSign, err error) {
	js.URL = url
	// npm install weixin-js-sdk
	ats := core.NewDefaultAccessTokenServer("wx267866e82ab809fc", "aa23d125ec35fd0724e12ec9352ce0d0", nil)
	clt := core.NewClient(ats, nil)

	jsClt := jssdk.NewDefaultTicketServer(clt)

	jsStr, err := jsClt.Ticket()
	// fmt.Println(jsStr)
	if err != nil {
		return
	}
	js.NonceStr = GetRandomSalt()
	js.Timestamp = strconv.Itoa(int(GetTimestamp()))
	signature := jssdk.WXConfigSign(jsStr, js.NonceStr, js.Timestamp, js.URL)
	// fmt.Println(signature)

	if signature == "" {
		return js, errors.New("创建令牌失败，或输入了不支持的域名")
	}

	js.Signature = signature

	return
	// panic(signature)
}

// 获取当前时间
func GetTimestamp() int64 {
	t := time.Now().Unix()
	return t
}

// return len=16  salt
func GetRandomSalt() string {
	return GetRandomString(16)
}

//生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
