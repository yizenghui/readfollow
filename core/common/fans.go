package common

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/user"
)

const wxAppID = "wx702b93aef72f3549"
const wxAppSecret = "8b69f45fc737a938cbaaffc05b192394"

// GetFans 关注通知
/*
	新追XXX
*/
func GetFans(OpenID string) (info *user.UserInfo, err error) {

	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)
	// user,err := user.Get(clt,OpenID,"zh_CN")
	return user.Get(clt, OpenID, "zh_CN")
}
