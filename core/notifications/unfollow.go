package notifications

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/message/template"
)

// Unfollow 取消关注书籍更新通知
/*
   弃追XXX
*/
func Unfollow(ToUser, Name string) (msgID int64, err error) {

	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)

	// UnfollowMSG 关注通知消息结构
	type UnfollowMSG struct {
		Name template.DataItem `json:"name"`
	}

	msg := template.TemplateMessage2{
		ToUser:     ToUser,
		TemplateId: "TDCLuaVwbKv_DjWGlMlZH79efrqqE330psg3A8yfflM",
		URL:        "",
		Data: UnfollowMSG{
			Name: template.DataItem{Value: Name, Color: ""},
		},
	}

	return template.Send(clt, msg)
}
