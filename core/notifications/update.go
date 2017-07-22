package notifications

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/message/template"
)

// UpdateMSG 消息结构
type UpdateMSG struct {
	Name    template.DataItem `json:"name"`
	Chapter template.DataItem `json:"chapter"`
}

// Update 更新通知
/*
   您关注的XXX已经更新到XXX章节了！
*/
func Update(ToUser, Name, Chapter, url string) (msgID int64, err error) {

	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)

	msg := template.TemplateMessage2{
		ToUser:     ToUser,
		TemplateId: "czFgqLwZ-39nMW_nAfBAs-U4ZnMAddj5uAY00oHG3cc",
		URL:        url,
		Data: UpdateMSG{
			Name:    template.DataItem{Value: Name, Color: "#173177"},
			Chapter: template.DataItem{Value: Chapter, Color: "#173177"},
		},
	}

	return template.Send(clt, msg)
}
