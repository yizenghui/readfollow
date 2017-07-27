package repository

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yizenghui/readfollow/conf"
	"github.com/yizenghui/readfollow/core/common"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"github.com/yizenghui/readfollow/model"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server

//	fansCache  *cache.Cache
)

func init() {

	conf.InitConfig("../conf/conf.toml")
	//	fansCache = cache.New(5*time.Minute, 30*time.Second)
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	msgHandler = mux

	// fmt.Println("xx", conf.Conf.Wechat.OriID, conf.Conf.Wechat.AppID, conf.Conf.Wechat.Token, conf.Conf.Wechat.AesKey)

	msgServer = core.NewServer(conf.Conf.Wechat.OriID, conf.Conf.Wechat.AppID, conf.Conf.Wechat.Token, conf.Conf.Wechat.AesKey, msgHandler, nil)
}

func textMsgHandler(ctx *core.Context) {

	// log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)

	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, "请多指教")

	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复

	//	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp) // 明文回复
	//	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	// log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {

	event := menu.GetClickEvent(ctx.MixedMsg)

	log.Println(event.EventKey)

	fans, _ := common.GetFans(event.FromUserName)
	// event.FromUserName

	openID := fans.OpenId

	var user model.User
	user.GetUserByOpenID(openID)

	if user.Nickname != fans.Nickname {
		user.Nickname = fans.Nickname
		user.Head = fans.HeadImageURL
		user.Save()
	}

	switch key := event.EventKey; key {

	case "myfollow":

		rc := fmt.Sprintf(`<a href="http://readfollow.com/u/%d?open_id=%v">%v的关注</a>`, user.ID, user.OpenID, user.Nickname)
		resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, rc)
		// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
		ctx.RawResponse(resp)

	case "new":

		rc := fmt.Sprintf(`<a href="http://readfollow.com/new?open_id=%v">最近更新</a>`, user.OpenID)
		resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, rc)
		// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
		ctx.RawResponse(resp)

	case "find":
		rc := fmt.Sprintf(`<a href="http://readfollow.com/find?open_id=%v">搜索书籍</a>`, user.OpenID)
		resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, rc)
		// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
		ctx.RawResponse(resp)

	default:
		resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "Please look forward to more features!")
		ctx.RawResponse(resp)
		// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	}

	//ctx.RawResponse(resp) // 明文回复
	//	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *core.Context) {

	event := menu.GetScanCodePushEvent(ctx.MixedMsg)

	SetWebGetSignTaskValueForWechatPush(event.EventKey, event.FromUserName)

	rc := fmt.Sprintf(`todo %v`, event.EventKey)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, rc)
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	ctx.RawResponse(resp)

	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

// WechatServe 微信接口服务
func WechatServe(w http.ResponseWriter, r *http.Request) {
	msgServer.ServeHTTP(w, r, nil)
}
