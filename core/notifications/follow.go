package notifications

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/message/template"
)

// Follow 关注通知
/*
	新追XXX
*/
func Follow(ToUser, Name string) (msgID int64, err error) {

	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)

	// FollowMSG 关注通知消息结构
	type FollowMSG struct {
		Name template.DataItem `json:"name"`
	}

	msg := template.TemplateMessage2{
		ToUser:     ToUser,
		TemplateId: "jN1NL9EJvG2bD8EtejiH6liVcxynfHhcl-AlwEwM-l0",
		URL:        "",
		Data: FollowMSG{
			Name: template.DataItem{Value: Name, Color: ""},
		},
	}

	return template.Send(clt, msg)
}
